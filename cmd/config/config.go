package config

import (
	"log"
	"net/url"
)

type AppConfig struct {
	Logger                  *log.Logger
	StorageURL              *url.URL
	StorageBucket           string
	PostgresConn            PostgresConn
	PasswordResetCodeExpiry int
	InProduction            bool
	TokenLifeTime           int
}

func NewAppConfig(
	infoLog *log.Logger,
	storageURL *url.URL,
	storageBucket string,
	passwordResetCodeExpiry int,
	inProduction bool,
	tokenLifeTime int,
	postgresConn PostgresConn,

) *AppConfig {
	return &AppConfig{
		Logger:                  infoLog,
		StorageURL:              storageURL,
		StorageBucket:           storageBucket,
		PasswordResetCodeExpiry: passwordResetCodeExpiry,
		InProduction:            inProduction,
		TokenLifeTime:           tokenLifeTime,
		PostgresConn:            postgresConn,
	}
}

type PostgresConn struct {
	URL    string
	DBName string
	DBUser string
	DBPass string
	Port   string
}
