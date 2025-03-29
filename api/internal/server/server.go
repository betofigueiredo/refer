package server

import (
	"refer/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uptrace/bun"
)

type Application struct {
	// UserHandler    *api.UserHandler
	// Middleware     middleware.UserMiddleware
	queries *bun.DB
	DBPool  *pgxpool.Pool
}

func NewApplication() (*Application, error) {
	dbpool, err := db.Open()
	if err != nil {
		return nil, err
	}

	queries, err := db.New(dbpool)
	if err != nil {
		panic(err)
	}

	// err = db.MigrateFS(pgDB, migrations.FS, ".")
	// if err != nil {
	// 	panic(err)
	// }

	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		// Logger:         logger,
		// UserHandler:    userHandler,
		// Middleware:     middlewareHandler,
		queries: queries,
		DBPool:  dbpool,
	}

	return app, nil
}
