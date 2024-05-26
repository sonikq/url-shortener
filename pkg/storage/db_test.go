package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"reflect"
	"testing"
	"time"
)

type testDB struct {
	dsn       string
	pool      *pgxpool.Pool
	container testcontainers.Container
}

func newTestDB() (db testDB, err error) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		return db, fmt.Errorf("failed to start container: %w", err)
	}

	ipAddress, err := postgresContainer.Host(ctx)
	if err != nil {
		return db, fmt.Errorf("failed to get container host: %w", err)
	}

	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		return db, fmt.Errorf("failed to get mapped port: %w", err)
	}

	dsn := fmt.Sprintf("postgres://user:password@%s:%s/testdb", ipAddress, mappedPort.Port())

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return db, fmt.Errorf("failed to parse config: %w", err)
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return db, fmt.Errorf("failed to connect to pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return db, fmt.Errorf("failed to ping pool: %w", err)
	}

	if err = dropTable(ctx, pool); err != nil {
		return db, err
	}

	if err = createTable(ctx, pool); err != nil {
		return db, err
	}

	if err = createIndex(ctx, pool); err != nil {
		return db, err
	}

	return testDB{
		dsn:       dsn,
		pool:      pool,
		container: postgresContainer,
	}, err
}

func (db *testDB) close() {
	if db != nil {
		db.pool.Close()
		_ = db.container.Terminate(context.Background())
	}
}

func Test_newDB(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)

	tests := []struct {
		name          string
		dsn           string
		dbPoolWorkers int
		wantErr       bool
	}{
		{
			name:          "valid dsn",
			dsn:           db.dsn,
			dbPoolWorkers: 0,
			wantErr:       false,
		},
		{
			name:          "invalid dsn",
			dsn:           "",
			dbPoolWorkers: 20,
			wantErr:       true,
		},
		{
			name:          "not existing database",
			dsn:           db.dsn + "non-existing",
			dbPoolWorkers: 8,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			var dbs *dbStorage
			dbs, err = newDB(ctx, tt.dsn, tt.dbPoolWorkers)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, dbs)
		})
	}
}

func Test_createTable(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid_table_name",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = createTable(context.Background(), db.pool)
			require.NoError(t, err)
		})
	}
}

func Test_createIndex(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid_index",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createIndex(context.Background(), db.pool); (err != nil) != tt.wantErr {
				t.Errorf("createIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dbStorage_Close(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	s := &dbStorage{
		pool: db.pool,
	}

	s.Close()
}

func Test_dbStorage_DeleteBatch(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	c := &dbStorage{
		pool: db.pool,
	}

	tests := []struct {
		name    string
		urls    []string
		userID  string
		wantErr bool
	}{

		{
			name:    "zero length urls",
			urls:    make([]string, 0),
			wantErr: false,
		},
		{
			name:    "normal urls",
			urls:    []string{"fjdks1fk", "16gdz0", "76fjsdh"},
			wantErr: false,
		},
		{
			name:    "valid",
			urls:    []string{"fjdks1fk", "16gdz0", "76fjsdh"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteBatch(context.Background(), tt.urls, tt.userID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dbStorage_Get(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	c := &dbStorage{
		pool: db.pool,
	}

	_ = c.Set(context.Background(), map[string]Item{
		"iuhpj21": {
			Object:     "https://yandex.ru",
			UserID:     "3pjojojngf",
			IsDeleted:  false,
			Expiration: 0,
		},
	})

	tests := []struct {
		name    string
		alias   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid-url",
			alias:   "iuhpj21",
			want:    "https://yandex.ru",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Get(context.Background(), tt.alias)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbStorage_GetBatchByUserID(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	c := &dbStorage{
		pool: db.pool,
	}

	_ = c.Set(context.Background(), map[string]Item{
		"iuhpj21": {
			Object:     "https://yandex.ru",
			UserID:     "3pjojojngf",
			IsDeleted:  false,
			Expiration: 0,
		},
	})

	tests := []struct {
		name    string
		userID  string
		want    map[string]Item
		wantErr bool
	}{
		{
			name:   "valid-batch",
			userID: "3pjojojngf",
			want: map[string]Item{
				"iuhpj21": {
					Object:     "https://yandex.ru",
					UserID:     "3pjojojngf",
					IsDeleted:  false,
					Expiration: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetBatchByUserID(context.Background(), tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBatchByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBatchByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbStorage_GetShortURL(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	c := &dbStorage{
		pool: db.pool,
	}

	tests := []struct {
		name        string
		originalURL string
		want        string
		wantErr     bool
	}{
		{
			name:        "valid-url",
			originalURL: "https://yandex.ru",
			want:        "http://localhost:8080/iuhpj21",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetShortURL(context.Background(), tt.originalURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetShortURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbStorage_Ping(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	s := &dbStorage{
		pool: db.pool,
	}

	err = s.Ping(context.Background())
	require.NoError(t, err)
}

func Test_dbStorage_Set(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	c := &dbStorage{
		pool: db.pool,
	}
	tests := []struct {
		name    string
		data    map[string]Item
		wantErr bool
	}{
		{
			name: "valid-data",
			data: map[string]Item{
				"iuhpj21": {
					Object:     "https://yandex.ru",
					UserID:     "3pjojojngf",
					IsDeleted:  false,
					Expiration: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.Set(context.Background(), tt.data); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dropTable(t *testing.T) {
	db, err := newTestDB()
	defer db.close()
	require.NoError(t, err)
	require.NotNil(t, db.pool)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err = dropTable(context.Background(), db.pool); (err != nil) != tt.wantErr {
				t.Errorf("dropTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
