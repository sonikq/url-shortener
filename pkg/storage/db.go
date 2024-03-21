package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type dbStorage struct {
	pool *pgxpool.Pool
}

func newDB(ctx context.Context, dsn string) (*dbStorage, error) {
	t1 := time.Now()
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
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

func createTable(ctx context.Context, pool *pgxpool.Pool) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, createTableQuery)
	if err != nil {
		if errRollBack := tx.Rollback(ctx); errRollBack != nil {
			return fmt.Errorf("exec error: %w; rollback error: %w", err, errRollBack)
		}
		return err
	}

	return tx.Commit(ctx)
}

func createIndex(ctx context.Context, pool *pgxpool.Pool) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, createIndexQuery)
	if err != nil {
		if errRollBack := tx.Rollback(ctx); errRollBack != nil {
			return fmt.Errorf("exec error: %w; rollback error: %w", err, errRollBack)
		}
		return err
	}

	return tx.Commit(ctx)
}

func (c *dbStorage) Set(ctx context.Context, data map[string]Item) (*string, error) {
	if len(data) == 0 {
		return nil, nil
	}

	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	for key, item := range data {
		_, err = tx.Exec(ctx, setNewValuesInDB, item.Object, key)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == pgerrcode.UniqueViolation {
					return &item.Object, ErrAlreadyExists
				}
			}
			if errRollBack := tx.Rollback(ctx); errRollBack != nil {
				return nil, fmt.Errorf("exec error: %w; rollback error: %w", err, errRollBack)
			}
			return nil, err
		}
	}

	return nil, tx.Commit(ctx)
}

func (c *dbStorage) Get(ctx context.Context, alias string) (string, error) {
	var originalURL string
	if err := c.pool.QueryRow(ctx, getShortURL, alias).Scan(&originalURL); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return originalURL, nil
}

func (c *dbStorage) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

func (c *dbStorage) Close() {
	c.pool.Close()
}
