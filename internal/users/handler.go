package users

import (
	"net/http"
	"transacta/internal/validation"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *Repository
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":   "validation_failed",
			"details": validation.FormatValidationError(err),
		})
		return
	}

	err := h.Repo.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.Repo.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch users"})
		return
	}

	c.JSON(200, users)
}
