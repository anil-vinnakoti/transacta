package account

import "time"

type Account struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}

type Transfer struct {
	ID            int       `json:"id"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
