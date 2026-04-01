package main

import (
	"log"
	"transacta/internal/account"
	"transacta/internal/db"
	"transacta/internal/middleware"
	"transacta/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Connect()
	db.RunMigrations(database)

	router := gin.New()

	// Middleware order (IMPORTANT)
	router.Use(gin.Recovery())         // panic safety
	router.Use(middleware.RequestID()) // request tracing
	router.Use(middleware.Logger())    // custom logging

	// Init Repositories
	userRepo := users.NewRepository(database)
	accountRepo := account.NewRepository(database)

	// Init Handlers
	usersHandler := users.NewHandler(userRepo)
	accountsHandler := account.NewHandler(accountRepo)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// GET
	router.GET("/users", usersHandler.GetUsers)
	router.GET("/accounts", accountsHandler.GetAccounts)
	router.GET("/transfers", accountsHandler.GetTransfers)

	// POST
	router.POST("/users", usersHandler.CreateUser)
	router.POST("/accounts", accountsHandler.CreateAccount)
	router.POST("/transfers", accountsHandler.Transfer)

	log.Println("🚀 Server running on :8080")
	router.Run(":8080")

}
