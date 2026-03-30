package main

import (
	"transacta/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Connect()

	db.RunMigrations(database)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")

}
