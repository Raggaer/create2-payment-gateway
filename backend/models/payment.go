package models

import (
	"database/sql"
	"time"
)

const (
	paymentStatusWaiting = iota
	paymentStatusDeployed
	paymentStatusFailed
)

// Payment is a struct that holds the payment information
type Payment struct {
	ID        int64
	Amount    int
	Status    int
	Address   string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreatePayment creates a new payment
func CreatePayment(salt string, amount int, addr string, db *sql.DB) (int64, error) {
	data, err := db.Exec(`
    INSERT INTO payments (salt, amount, address, status, created_at, updated_at)
    VALUES (?, ?, ?, ?, NOW(), NOW())
  `, salt, amount, addr, paymentStatusWaiting)
	if err != nil {
		return 0, err
	}
	return data.LastInsertId()
}

// RetrieveWaitingPayments retrieves all payments that are waiting for funds to arrive
func RetrieveWaitingPayments(db *sql.DB) ([]Payment, error) {
	rows, err := db.Query(`
    SELECT id, salt, amount, address, status, created_at, updated_at
    FROM payments
    WHERE status = ?
  `, paymentStatusWaiting)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var p Payment
		err := rows.Scan(&p.ID, &p.Salt, &p.Amount, &p.Address, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		payments = append(payments, p)
	}
	return payments, nil
}
