package server

import (
	"database/sql"
	"refer/internal/api"
	"refer/internal/db"

	"github.com/uptrace/bun"
)

type Application struct {
	// Middleware     middleware.UserMiddleware
	UserHandler *api.UserHandler
	Queries     *bun.DB
	SQLDB       *sql.DB
}

func NewApplication() (*Application, error) {
	sqldb, err := db.Open()
	if err != nil {
		return nil, err
	}

	queries, err := db.New(sqldb)
	if err != nil {
		panic(err)
	}

	// err = db.MigrateFS(pgDB, migrations.FS, ".")
	// if err != nil {
	// 	panic(err)
	// }

	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// handlers
	userHandler := api.NewUsersHandler(queries)

	app := &Application{
		// Logger:         logger,
		// Middleware:     middlewareHandler,
		UserHandler: userHandler,
		Queries:     queries,
		SQLDB:       sqldb,
	}

	return app, nil
}
