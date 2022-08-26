// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: payments.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgtype"
)

const createPaymentTransaction = `-- name: CreatePaymentTransaction :one

insert into payment_transactions (created_at,
                                  updated_at,
                                  transaction_ref,
                                  status,
                                  transaction_complete,
                                  data,
                                  code,
                                  payment_integration_type,
                                  payment_purpose,
                                  amount,
                                  payment_mode,
                                  currency, phone_number)
VALUES (current_timestamp,
        current_timestamp,
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11)
returning id
`

type CreatePaymentTransactionParams struct {
	TransactionRef         string       `json:"transaction_ref"`
	Status                 bool         `json:"status"`
	TransactionComplete    bool         `json:"transaction_complete"`
	Data                   pgtype.JSONB `json:"data"`
	Code                   string       `json:"code"`
	PaymentIntegrationType int32        `json:"payment_integration_type"`
	PaymentPurpose         int32        `json:"payment_purpose"`
	Amount                 float64      `json:"amount"`
	PaymentMode            int32        `json:"payment_mode"`
	Currency               string       `json:"currency"`
	PhoneNumber            string       `json:"phone_number"`
}

// TODO:Add indexing for paymentransactions
func (q *Queries) CreatePaymentTransaction(ctx context.Context, arg CreatePaymentTransactionParams) (int64, error) {
	row := q.db.QueryRow(ctx, createPaymentTransaction,
		arg.TransactionRef,
		arg.Status,
		arg.TransactionComplete,
		arg.Data,
		arg.Code,
		arg.PaymentIntegrationType,
		arg.PaymentPurpose,
		arg.Amount,
		arg.PaymentMode,
		arg.Currency,
		arg.PhoneNumber,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getPaymentTransactionByID = `-- name: GetPaymentTransactionByID :one
select payment_transactions.id,
       created_at,
       updated_at,
       transaction_ref,
       status,
       transaction_complete,
       data,
       code,
       payment_integration_type,
       payment_purpose,
       amount,
       payment_mode,
       currency,
       phone_number
from payment_transactions
where id = $1
`

type GetPaymentTransactionByIDRow struct {
	ID                     int64        `json:"id"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
	TransactionRef         string       `json:"transaction_ref"`
	Status                 bool         `json:"status"`
	TransactionComplete    bool         `json:"transaction_complete"`
	Data                   pgtype.JSONB `json:"data"`
	Code                   string       `json:"code"`
	PaymentIntegrationType int32        `json:"payment_integration_type"`
	PaymentPurpose         int32        `json:"payment_purpose"`
	Amount                 float64      `json:"amount"`
	PaymentMode            int32        `json:"payment_mode"`
	Currency               string       `json:"currency"`
	PhoneNumber            string       `json:"phone_number"`
}

func (q *Queries) GetPaymentTransactionByID(ctx context.Context, id int64) (*GetPaymentTransactionByIDRow, error) {
	row := q.db.QueryRow(ctx, getPaymentTransactionByID, id)
	var i GetPaymentTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TransactionRef,
		&i.Status,
		&i.TransactionComplete,
		&i.Data,
		&i.Code,
		&i.PaymentIntegrationType,
		&i.PaymentPurpose,
		&i.Amount,
		&i.PaymentMode,
		&i.Currency,
		&i.PhoneNumber,
	)
	return &i, err
}

const getTotalPayments = `-- name: GetTotalPayments :one
SELECT COALESCE(SUM(amount), 0) AS total
FROM payment_transactions
`

func (q *Queries) GetTotalPayments(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRow(ctx, getTotalPayments)
	var total interface{}
	err := row.Scan(&total)
	return total, err
}

const listPaymentTransactions = `-- name: ListPaymentTransactions :many
select payment_transactions.id,
       created_at,
       updated_at,
       transaction_ref,
       status,
       transaction_complete,
       data,
       code,
       payment_integration_type,
       payment_purpose,
       amount,
       payment_mode,
       currency,
       phone_number
from payment_transactions
`

type ListPaymentTransactionsRow struct {
	ID                     int64        `json:"id"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
	TransactionRef         string       `json:"transaction_ref"`
	Status                 bool         `json:"status"`
	TransactionComplete    bool         `json:"transaction_complete"`
	Data                   pgtype.JSONB `json:"data"`
	Code                   string       `json:"code"`
	PaymentIntegrationType int32        `json:"payment_integration_type"`
	PaymentPurpose         int32        `json:"payment_purpose"`
	Amount                 float64      `json:"amount"`
	PaymentMode            int32        `json:"payment_mode"`
	Currency               string       `json:"currency"`
	PhoneNumber            string       `json:"phone_number"`
}

func (q *Queries) ListPaymentTransactions(ctx context.Context) ([]*ListPaymentTransactionsRow, error) {
	rows, err := q.db.Query(ctx, listPaymentTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ListPaymentTransactionsRow{}
	for rows.Next() {
		var i ListPaymentTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TransactionRef,
			&i.Status,
			&i.TransactionComplete,
			&i.Data,
			&i.Code,
			&i.PaymentIntegrationType,
			&i.PaymentPurpose,
			&i.Amount,
			&i.PaymentMode,
			&i.Currency,
			&i.PhoneNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPaymentTransactionsByUser = `-- name: ListPaymentTransactionsByUser :many
select payment_transactions.id,
       created_at,
       updated_at,
       transaction_ref,
       status,
       transaction_complete,
       data,
       code,
       payment_integration_type,
       payment_purpose,
       amount,
       payment_mode,
       currency,
       phone_number
from payment_transactions
where phone_number = $1
`

type ListPaymentTransactionsByUserRow struct {
	ID                     int64        `json:"id"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
	TransactionRef         string       `json:"transaction_ref"`
	Status                 bool         `json:"status"`
	TransactionComplete    bool         `json:"transaction_complete"`
	Data                   pgtype.JSONB `json:"data"`
	Code                   string       `json:"code"`
	PaymentIntegrationType int32        `json:"payment_integration_type"`
	PaymentPurpose         int32        `json:"payment_purpose"`
	Amount                 float64      `json:"amount"`
	PaymentMode            int32        `json:"payment_mode"`
	Currency               string       `json:"currency"`
	PhoneNumber            string       `json:"phone_number"`
}

func (q *Queries) ListPaymentTransactionsByUser(ctx context.Context, phoneNumber string) ([]*ListPaymentTransactionsByUserRow, error) {
	rows, err := q.db.Query(ctx, listPaymentTransactionsByUser, phoneNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ListPaymentTransactionsByUserRow{}
	for rows.Next() {
		var i ListPaymentTransactionsByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TransactionRef,
			&i.Status,
			&i.TransactionComplete,
			&i.Data,
			&i.Code,
			&i.PaymentIntegrationType,
			&i.PaymentPurpose,
			&i.Amount,
			&i.PaymentMode,
			&i.Currency,
			&i.PhoneNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}