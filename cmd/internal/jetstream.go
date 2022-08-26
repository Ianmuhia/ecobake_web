package internal

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func GetJetStream(natsURL string, logger *log.Logger) nats.JetStreamContext {
	if natsURL != "" {
		natsURL = nats.DefaultURL
	}
	nc, err := nats.Connect(natsURL,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second))
	if err != nil {
		logger.Println(err)
		return nil
	}
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		logger.Println(err)
		return nil
	}
	logger.Printf("nats has connected ...%v", nc.ConnectedAddr())
	return js
}
