package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/1ef7yy/effective_mobile_test/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	log logger.Logger
	DB  *pgxpool.Pool
}

func Config(dsn string, log logger.Logger) *pgxpool.Config {
	const defaultMaxConns = 10
	const defaultMinConns = 0
	const defaultMaxConnLifetime = time.Hour * 1
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	dbConfig, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create a config: %s", err))
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	log.Info(fmt.Sprintf("postgres: defaultMaxConns: %d", dbConfig.MaxConns))
	log.Info(fmt.Sprintf("postgres: defaultMinConns: %d", dbConfig.MinConns))
	log.Info(fmt.Sprintf("postgres: defaultMaxConnLifetime: %s", dbConfig.MaxConnLifetime))
	log.Info(fmt.Sprintf("postgres: defaultMaxConnIdleTime: %s", dbConfig.MaxConnIdleTime))
	log.Info(fmt.Sprintf("postgres: defaultHealthCheckPeriod: %s", dbConfig.HealthCheckPeriod))
	log.Info(fmt.Sprintf("postgres: defaultConnectTimeout: %s", dbConfig.ConnConfig.ConnectTimeout))

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Info("Closed the connection pool.")
	}

	return dbConfig
}

func NewPostgres(ctx context.Context, dsn string, log logger.Logger) (*Postgres, error) {
	var (
		pgInstance *Postgres
		pgOnce     sync.Once
		pgErr      error
	)

	pgOnce.Do(func() {
		db, err := pgxpool.NewWithConfig(ctx, Config(dsn, log))
		if err != nil {
			log.Fatal("Unable to connect to database: " + err.Error())
			pgErr = err
		}

		pgInstance = &Postgres{
			log: log,
			DB:  db,
		}
	})

	if pgErr != nil {
		return nil, pgErr
	}
	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}
