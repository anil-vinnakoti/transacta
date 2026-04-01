package account

import (
	"net/http"
	"transacta/internal/validation"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *Repository
}

type CreateAccountRequest struct {
	UserID  int     `json:"user_id" binding:"required,gt=0"`
	Balance float64 `json:"balance" binding:"required,gte=0"`
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":   "validation_failed",
			"details": validation.FormatValidationError(err),
		})
		return
	}

	id, err := h.Repo.Create((req.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account_id": id})
}

func (h *Handler) Transfer(c *gin.Context) {
	var req TransferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":   "validation_failed",
			"details": validation.FormatValidationError(err),
		})
		return
	}

	err := h.Repo.Transfer(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *Handler) GetAccounts(c *gin.Context) {
	accounts, err := h.Repo.GetAccounts()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch accounts"})
		return
	}
	c.JSON(200, accounts)
}

func (h *Handler) GetTransfers(c *gin.Context) {
	transfers, err := h.Repo.GetTransfers()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch transactions"})
		return
	}
	c.JSON(200, transfers)
}
