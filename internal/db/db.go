package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	var db *sql.DB
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("✅ Connected to DB")
				return db
			}
		}

		log.Println("⏳ Waiting for DB...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Could not connect to DB:", err)
	return nil
}
