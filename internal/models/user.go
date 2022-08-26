package models

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64     `json:"id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"-"`
	UserName     string    `json:"user_name,omitempty"`
	Email        string    `json:"email,omitempty"`
	PasswordHash string    `json:"password_hash,"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"password"`
	ProfileImage string    `json:"profile_image,"`
	IsVerified   bool      `json:"is_verified,omitempty"`
}

type Geo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (u *User) Hash() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	u.PasswordHash = string(bytes)
	return err
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) ComparePassword(passwordConfirm string) bool {
	return u.Password == passwordConfirm
}

func (u *User) FormatDate(time2 string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, time2)
	u.CreatedAt = t
	return t
}

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
