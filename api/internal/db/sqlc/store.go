package db

import (
	"fmt"
	"runtime"

	"database/sql"

	"github.com/uptrace/bun/driver/sqliteshim"
)

func Open() (*sql.DB, error) {
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	// sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	sqldb, err := sql.Open(sqliteshim.ShimName, "/refer.db")
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}
	maxOpenConns := 4 * runtime.GOMAXPROCS(0)

	// pool.SetConnMaxLifetime(0)
	// pool.SetMaxIdleConns(3)
	// pool.SetMaxOpenConns(3)

	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)
	return sqldb, nil

	// dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_ADDR"))
	// if err != nil {
	// 	return nil, fmt.Errorf("db: open %w", err)
	// }
	// return dbpool, nil
}

// func New2(sqldb *sql.DB) (*bun.DB, error) {
// 	db := bun.NewDB(sqldb, sqlitedialect.New())
// 	// sqldb := stdlib.OpenDBFromPool(dbpool)
// 	// db := bun.NewDB(sqldb, pgdialect.New())
// 	return db, nil
// }
