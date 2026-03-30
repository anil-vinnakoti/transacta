package db

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE
		);`,
		`CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			balance NUMERIC DEFAULT 0
		);`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatal("❌ Migration failed:", err)
		}
	}

	log.Println("✅ Migrations ran successfully")

}
