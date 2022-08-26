package config

import (
	"log"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
)

type AppConfig struct {
	Logger                  *log.Logger
	StorageURL              *url.URL
	StorageBucket           string
	PostgresConn            PostgresConn
	PasswordResetCodeExpiry int
	InProduction            bool
	TokenLifeTime           int
	RedisConn               *redis.Client
	Js                      nats.JetStreamContext
	ZincChan                chan any
	ZincRcvChan             chan any
}

func NewAppConfig(
	infoLog *log.Logger,
	storageURL *url.URL,
	storageBucket string,
	passwordResetCodeExpiry int,
	inProduction bool,
	tokenLifeTime int,
	redisConn *redis.Client,
	js nats.JetStreamContext,
	postgresConn PostgresConn,
	zincChan chan any,
	zincRcvChan chan any,

) *AppConfig {
	return &AppConfig{
		Logger:                  infoLog,
		StorageURL:              storageURL,
		StorageBucket:           storageBucket,
		PasswordResetCodeExpiry: passwordResetCodeExpiry,
		InProduction:            inProduction,
		TokenLifeTime:           tokenLifeTime,
		RedisConn:               redisConn,
		Js:                      js,
		PostgresConn:            postgresConn,
		ZincChan:                zincChan,
		ZincRcvChan:             zincRcvChan,
	}
}

type PostgresConn struct {
	URL    string
	DBName string
	DBUser string
	DBPass string
	Port   string
}

type ZincIndexes struct {
	Users      string
	Ventures   string
	Businesses string
}
