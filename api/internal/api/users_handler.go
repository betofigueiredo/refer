package api

import (
	"refer/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type UserHandler struct {
	store *bun.DB
}

func NewUsersHandler(store *bun.DB) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []db.User
	user := new(db.User)
	err := h.store.NewSelect().Model(user).Scan(c, &users)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"users": users})
}
