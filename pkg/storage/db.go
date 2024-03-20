package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type postgresClient struct {
	pool *pgxpool.Pool
}

func newPostgresClient(ctx context.Context, dsn string) (*postgresClient, error) {
	t1 := time.Now()
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	//log.Println(config)

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	log.Printf("connection to database took: %v\n", time.Since(t1))

	return &postgresClient{pool: pool}, nil
}

func (c *postgresClient) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

func (c *postgresClient) Close() {
	c.pool.Close()
}
