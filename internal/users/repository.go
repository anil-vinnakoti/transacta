package users

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(u *CreateUserRequest) error {
	query := `INSERT INTO users (name, email) values ($1, $2) RETURNING id`

	return r.DB.QueryRow(query, u.Name, u.Email).Scan()
}

func (r *Repository) GetUsers() ([]User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email from users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
