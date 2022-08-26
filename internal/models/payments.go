package models

import "time"

type PaymentsTransaction struct {
	ID                     int64       `json:"id"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	TransactionRef         string      `json:"transaction_ref"`
	Status                 bool        `json:"status"`
	TransactionComplete    bool        `json:"transaction_complete"`
	Data                   interface{} `json:"data"`
	Code                   string      `json:"code"`
	PaymentIntegrationType int32       `json:"payment_integration_type"`
	PaymentPurpose         int32       `json:"payment_purpose"`
	Amount                 int64       `json:"amount"`
	PaymentMode            int32       `json:"payment_mode"`
	Currency               string      `json:"currency"`
	PhoneNumber            string      `json:"phone_number"`
}
