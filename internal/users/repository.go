package users

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(u *User) error {
	query := `INSERT INTO users (name, email) values ($1, $2) RETURNING id`

	return r.DB.QueryRow(query, u.Name, u.Email).Scan(&u.ID)
}
