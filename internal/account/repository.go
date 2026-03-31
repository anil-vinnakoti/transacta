package account

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(userID int) (int, error) {
	var id int
	query := `INSERT INTO accounts (user_id, balance) VALUES ($1, 0) RETURNING id`

	err := r.DB.QueryRow(query, userID).Scan(&id)
	return id, err
}
