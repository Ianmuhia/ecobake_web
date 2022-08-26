package services

import (
	"context"
	"ecobake/cmd/config"
	"log"
	"net/url"
	"time"
)

type mailService struct {
	cfg *config.AppConfig
}

type MailService interface {
	VerifyMailCode(key string) string
	RemoveMailCode(key string)
	NewMail(from string, to string, subject string, mailType MailType, body *MailData) *Mail
}

func NewMailService() MailService {
	return &mailService{}
}

type MailType int

const (
	MailConfirmation MailType = iota + 1
	PassReset
)

// MailData represents the data to be sent to the template of the mail.
type MailData struct {
	Username string
	Code     string
	URL      *url.URL
}

// Mail represents a email request.
type Mail struct {
	From     string
	To       string
	Subject  string
	Body     *MailData
	MailType MailType
}

// VerificationData represents the type for the data stored for verification.
type VerificationData struct {
	Email     string    `json:"email" validate:"required" sql:"email"`
	Code      string    `json:"code" validate:"required" sql:"code"`
	ExpiresAt time.Time `json:"expires_at" `
	Type      string    `json:"type" sql:"type"`
}

func (ms *mailService) NewMail(from string, to string, subject string, mailType MailType, body *MailData) *Mail {
	return &Mail{
		From:     from,
		To:       to,
		Subject:  subject,
		MailType: mailType,
		Body:     body,
	}
}

func (ms *mailService) VerifyMailCode(key string) string {
	data, err := ms.cfg.RedisConn.Get(context.TODO(), key).Result()
	log.Printf("Redis Code %v or %v", data, err)
	if err != nil {
		log.Println(err)
		return "Invalid key provided or key not found"
	}
	return "data"
}

func (ms *mailService) RemoveMailCode(key string) {
	var foundedRecordCount = 0
	iter := ms.cfg.RedisConn.Scan(context.Background(), 0, key, 0).Iterator()
	for iter.Next(context.Background()) {
		ms.cfg.RedisConn.Del(context.Background(), iter.Val())
		foundedRecordCount++
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}
	err := ms.cfg.RedisConn.Close()
	if err != nil {
		log.Println(err)
		return
	}
}
