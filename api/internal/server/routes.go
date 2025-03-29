package server

import (
	"net/http"
	"refer/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func SetupRoutes(app *Application) http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/test", func(c *gin.Context) {
		uuid1 := "056dbf4f-0190-43cc-b3e4-9828201d002b"
		by := uuid.MustParse(uuid1)
		u, _ := pgtype.UUID{Bytes: by, Valid: true}.UUIDValue()
		user := new(db.User)
		err := app.queries.NewSelect().Model(user).Where("id = ?", 1).Scan(c)
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err})
			return
		}
		c.JSON(200, gin.H{"message": "ponssg", "result": u})
	})

	return r
}
