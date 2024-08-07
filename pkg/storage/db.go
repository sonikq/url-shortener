package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sonikq/url-shortener/internal/app/models"
	"log"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type dbStorage struct {
	pool *pgxpool.Pool
}

func newDB(ctx context.Context, dsn string, dbPoolWorkers int) (*dbStorage, error) {
	t1 := time.Now()
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	config.MaxConns = int32(dbPoolWorkers)

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err = dropTable(ctx, pool); err != nil {
		return nil, err
	}

	if err = createTable(ctx, pool); err != nil {
		return nil, err
	}

	if err = createIndex(ctx, pool); err != nil {
		return nil, err
	}

	log.Printf("connection to database took: %v\n", time.Since(t1))

	return &dbStorage{pool: pool}, nil
}

func dropTable(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, preCreateTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func createTable(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func createIndex(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, createOriginalURLIndexQuery)
	if err != nil {
		return err
	}

	return nil
}

// Set -
func (c *dbStorage) Set(ctx context.Context, data map[string]Item) error {
	if len(data) == 0 {
		return nil
	}

	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error while begin transaction: %w", err)
	}
	defer func() {
		if errRollBack := tx.Rollback(ctx); errRollBack != nil {
			fmt.Printf("rollback error: %v", errRollBack)
		}
	}()

	for key, item := range data {
		_, err = tx.Exec(ctx, setNewValueInDB, item.Object, key, item.UserID)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == pgerrcode.UniqueViolation {
					return models.ErrAlreadyExists
				}
			}
			return err
		}
	}

	return tx.Commit(ctx)
}

// DeleteBatch -
func (c *dbStorage) DeleteBatch(ctx context.Context, urls []string, userID string) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {

		return fmt.Errorf("error while begin transaction: %s", err.Error())
	}
	defer func() {
		if errRollBack := tx.Rollback(ctx); errRollBack != nil {
			fmt.Printf("rollback error: %v", errRollBack)
		}
	}()

	for _, value := range urls {
		_, err = tx.Exec(ctx, setDeleteBatch, value, userID)
		if err != nil {
			return fmt.Errorf("cant execute db command: %s", err.Error())
		}
	}
	return tx.Commit(ctx)
}

// GetBatchByUserID -
func (c *dbStorage) GetBatchByUserID(ctx context.Context, userID string) (map[string]Item, error) {
	batch := make(map[string]Item)

	rows, err := c.pool.Query(ctx, getBatchByUserID, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var originalURL, shortURL string
		err = rows.Scan(&originalURL, &shortURL)
		if err != nil {
			return nil, err
		}
		batch[shortURL] = Item{
			Object: originalURL,
			UserID: userID,
		}
	}
	return batch, nil
}

// Get -
func (c *dbStorage) Get(ctx context.Context, alias string) (string, error) {
	var originalURL string
	var isDeleted bool
	if err := c.pool.QueryRow(ctx, getOriginalURL, alias).Scan(&originalURL, &isDeleted); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	if isDeleted {
		return "", models.ErrGetDeletedLink
	}

	return originalURL, nil
}

// GetShortURL -
func (c *dbStorage) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	var shortURL string
	if err := c.pool.QueryRow(ctx, getShortURL, originalURL).Scan(&shortURL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return shortURL, nil
}

// GetStats - returning distinct urls and users from storage
func (c *dbStorage) GetStats(ctx context.Context) (int64, int64, error) {
	var countOfURLs, countOfUsers sql.NullInt64
	if err := c.pool.QueryRow(ctx, getCountOfURLs).Scan(&countOfURLs); err != nil {
		return -1, -1, err
	}

	if err := c.pool.QueryRow(ctx, getCountOfUsers).Scan(&countOfUsers); err != nil {
		return -1, -1, err
	}
	return countOfURLs.Int64, countOfUsers.Int64, nil
}

// Ping -
func (c *dbStorage) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

// Close -
func (c *dbStorage) Close() {
	c.pool.Close()
}
