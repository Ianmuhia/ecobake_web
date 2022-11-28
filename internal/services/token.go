package services

import (
	"ecobake/internal/models"
	"errors"
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v4"
)

const minSecretKeySize = 5

type TokenService interface {
	CreateToken(username string, email string, duration time.Duration, rtduration time.Duration, id int64) (string, error)
	CreateRefreshToken(username string, email string, duration time.Duration, rtduration time.Duration, id int64) (string, error)
	VerifyToken(token string) (*models.Payload, error)
}

type tokenService struct {
	secretKey string
}

func NewTokenService(secretKey string) (*tokenService, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &tokenService{secretKey: secretKey}, nil
}

// Different types of error returned by the VerifyToken function.
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// NewPayload creates a new token payload with a specific username and duration.
func NewPayload(username string, email string, duration time.Duration, rtduration time.Duration, id int64) (*models.Payload, error) {
	payload := &models.Payload{
		ID:          id,
		Username:    username,
		Email:       email,
		IssuedAt:    time.Now(),
		ATExpiredAt: time.Now().Add(duration),
		RTExpiredAt: time.Now().Add(rtduration),
	}
	return payload, nil
}

// CreateToken creates a new token for a specific username and duration.
func (maker *tokenService) CreateToken(username string, email string, duration time.Duration, rtduration time.Duration, id int64) (string, error) {
	payload, err := NewPayload(username, email, duration, rtduration, id)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not.
func (maker *tokenService) VerifyToken(token string) (*models.Payload, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &models.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)

		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*models.Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

// CreateRefreshToken creates a new token for a specific email and duration.
func (maker *tokenService) CreateRefreshToken(username string, email string, duration time.Duration, rtduration time.Duration, id int64) (string, error) {
	payload, err := NewPayload(username, email, duration, rtduration, id)
	if err != nil {
		return "", err
	}

	refresherToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return refresherToken.SignedString([]byte(maker.secretKey))
}
