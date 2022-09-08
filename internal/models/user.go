package models

import (
	"errors"
	"time"
)

type Payload struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	IssuedAt    time.Time `json:"issued_at"`
	ATExpiredAt time.Time `json:"expired_at"`
	RTExpiredAt time.Time `json:"rt_expired_at"`
}

// Valid checks if the token payload is valid or not.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ATExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}
