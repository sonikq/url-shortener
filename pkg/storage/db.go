package storage

import (
	"context"
	"database/sql"
	"fmt"
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

func (c *dbStorage) Set(ctx context.Context, data map[string]Item) error {
	if len(data) == 0 {
		return nil
	}

	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}

	for key, item := range data {
		_, err = tx.Exec(ctx, setNewValuesInDB, item.Object, key)
		if err != nil {
			if errRollBack := tx.Rollback(ctx); errRollBack != nil {
				return fmt.Errorf("exec error: %w; rollback error: %w", err, errRollBack)
			}
			return err
		}
	}

	return tx.Commit(ctx)
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
