package server

import (
	"database/sql"
	"refer/internal/api"
	db "refer/internal/db/sqlc"
)

type Application struct {
	// Middleware     middleware.UserMiddleware
	UserHandler *api.UserHandler
	Queries     *db.Queries
	SQLDB       *sql.DB
}

func NewApplication() (*Application, error) {
	sqldb, err := db.Open()
	if err != nil {
		return nil, err
	}

	queries := db.New(sqldb)

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
