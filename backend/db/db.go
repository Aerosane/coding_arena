package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect(ctx context.Context, databaseURL string) error {
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return fmt.Errorf("parse database url: %w", err)
	}
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	Pool = pool
	log.Printf("[INFO] Connected to Postgres (max_conns=%d)", cfg.MaxConns)
	return nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}

func Migrate(ctx context.Context) error {
	_, err := Pool.Exec(ctx, schema)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}

const schema = `
CREATE TABLE IF NOT EXISTS submissions (
    id          TEXT PRIMARY KEY,
    problem_id  TEXT NOT NULL,
    language    TEXT NOT NULL,
    source      TEXT NOT NULL,
    verdict     TEXT NOT NULL DEFAULT 'pending',
    points      DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_time  DOUBLE PRECISION NOT NULL DEFAULT 0,
    max_memory  BIGINT NOT NULL DEFAULT 0,
    compile_error TEXT NOT NULL DEFAULT '',
    cases       JSONB NOT NULL DEFAULT '[]',
    ip          TEXT NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_submissions_problem ON submissions(problem_id);
CREATE INDEX IF NOT EXISTS idx_submissions_verdict ON submissions(verdict);
CREATE INDEX IF NOT EXISTS idx_submissions_created ON submissions(created_at DESC);
`
