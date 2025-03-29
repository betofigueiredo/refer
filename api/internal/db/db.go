package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Open() (*pgxpool.Pool, error) {
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_ADDR"))
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}
	return dbpool, nil
}

func New(dbpool *pgxpool.Pool) (*bun.DB, error) {
	sqldb := stdlib.OpenDBFromPool(dbpool)
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}
