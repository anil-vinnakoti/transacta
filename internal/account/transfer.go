package account

import (
	"errors"
)

type TransferRequest struct {
	FromAccountID int     `json:"from_account_id"`
	ToAccountID   int     `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

func (r *Repository) Transfer(req TransferRequest) error {
	// ✅ Validation
	if req.FromAccountID == req.ToAccountID {
		return errors.New("cannot transfer to same account")
	}

	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	// Rollback safety
	defer tx.Rollback()

	// Deadlock prevention: lock in order
	fromID := req.FromAccountID
	toID := req.ToAccountID

	if fromID > toID {
		fromID, toID = toID, fromID
	}

	var dummy float64

	// Lock first account
	err = tx.QueryRow(
		`SELECT balance FROM accounts WHERE id=$1 FOR UPDATE`,
		fromID,
	).Scan(&dummy)
	if err != nil {
		return err
	}

	// Lock second account
	err = tx.QueryRow(
		`SELECT balance FROM accounts WHERE id=$1 FOR UPDATE`,
		toID,
	).Scan(&dummy)
	if err != nil {
		return err
	}

	// Check sender balance AFTER locking
	var balance float64
	err = tx.QueryRow(
		`SELECT balance FROM accounts WHERE id=$1`,
		req.FromAccountID,
	).Scan(&balance)
	if err != nil {
		return err
	}

	if balance < req.Amount {
		return errors.New("insufficient funds")
	}

	// Deduct from sender
	_, err = tx.Exec(
		`UPDATE accounts SET balance = balance - $1 WHERE id=$2`,
		req.Amount,
		req.FromAccountID,
	)
	if err != nil {
		return err
	}

	// Add to receiver
	_, err = tx.Exec(
		`UPDATE accounts SET balance = balance + $1 WHERE id=$2`,
		req.Amount,
		req.ToAccountID,
	)
	if err != nil {
		return err
	}

	// Record transfer
	_, err = tx.Exec(
		`INSERT INTO transfers (from_account_id, to_account_id, amount)
		 VALUES ($1, $2, $3)`,
		req.FromAccountID,
		req.ToAccountID,
		req.Amount,
	)
	if err != nil {
		return err
	}

	// Commit transaction
	return tx.Commit()
}
