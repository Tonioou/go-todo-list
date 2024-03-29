package client

import (
	"context"
	"fmt"
	"github.com/Tonioou/go-todo-list/internal/config"
	logger "github.com/Tonioou/go-todo-list/pkg"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joomcode/errorx"
)

type PgClient struct {
	pgxCfg  *pgxpool.Config
	pgxConn *pgxpool.Pool
}

func NewPgClient(ctx context.Context, cfg *config.Replica) *PgClient {
	pgxCfg, err := pgxpool.ParseConfig(createConnString(cfg))
	if err != nil {
		logger.Fatal("cannot create pg config", err)
	}
	pgxCfg.ConnConfig.Tracer = otelpgx.NewTracer()
	connPool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		logger.Fatal("cannot connect to pg", err)
	}
	return &PgClient{
		pgxCfg:  pgxCfg,
		pgxConn: connPool,
	}
}

func createConnString(cfg *config.Replica) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database.Name,
	)
}

func (pg *PgClient) Ping(ctx context.Context) error {
	var (
		errx        *errorx.Error
		query       = "SELECT 1;"
		queryResult = 0
	)

	err := pg.pgxConn.QueryRow(ctx, query).Scan(&queryResult)
	if err != nil {
		errx = errorx.Decorate(err, "failed to query database")
	}
	return errx
}

func (pg *PgClient) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := pg.pgxConn.Query(ctx, query, args)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to query database")
	}
	return rows, nil
}

func (pg *PgClient) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := pg.pgxConn.Exec(ctx, query, args...)
	if err != nil {
		return errorx.Decorate(err, "failed to insert on database")
	}
	return nil
}

func (pg *PgClient) QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, error) {
	rows := pg.pgxConn.QueryRow(ctx, query, args...)
	return rows, nil
}
