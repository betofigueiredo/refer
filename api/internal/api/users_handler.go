package api

import (
	db "refer/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	store *db.Queries
}

func NewUsersHandler(store *db.Queries) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.store.GetUsers(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"users": users})
}
