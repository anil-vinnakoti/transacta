package main

import (
	"log"
	"transacta/internal/account"
	"transacta/internal/db"
	"transacta/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Connect()
	db.RunMigrations(database)

	router := gin.Default()

	// Init Repositories
	userRepo := users.NewRepository(database)
	accountRepo := account.NewRepository(database)

	// Init Handlers
	usersHandler := users.NewHandler(userRepo)
	accountsHandler := account.NewHandler(accountRepo)

	// Routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.POST("/users", usersHandler.CreateUser)
	router.POST("/accounts", accountsHandler.CreateAccount)
	router.POST("/transfer", accountsHandler.Transfer)

	log.Println("🚀 Server running on :8080")
	router.Run(":8080")

}
