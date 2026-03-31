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

func (r *Repository) GetAccounts() ([]Account, error) {
	rows, err := r.DB.Query(`SELECT id, user_id, balance FROM accounts`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var a Account
		err := rows.Scan(&a.ID, &a.UserID, &a.Balance)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}
