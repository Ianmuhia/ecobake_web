package services

import (
	"ecobake/cmd/config"
	"github.com/nats-io/nats.go"
	"log"
)

const (
	streamSubjects = "Mail.*"
)

type NatsService interface {
	CreateStream(streamDescription string, streamName string) error
	Publish(subjectName string, data []byte) error
}

type natsService struct {
	cfg *config.AppConfig
}

func NewNatsService(cfg *config.AppConfig) NatsService {
	return &natsService{cfg: cfg}
}

func (ns *natsService) CreateStream(streamDescription string, streamName string) error {
	// Check if the streamName stream already exists; if not, create it.
	stream, err := ns.cfg.Js.StreamInfo(streamName)
	if err != nil {
		ns.cfg.Logger.Println(err)
		return err
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = ns.cfg.Js.AddStream(&nats.StreamConfig{
			Name:              streamName,
			Description:       streamDescription,
			Subjects:          []string{streamSubjects},
			Retention:         0,
			MaxConsumers:      0,
			MaxMsgs:           0,
			MaxBytes:          0,
			Discard:           0,
			MaxAge:            0,
			MaxMsgsPerSubject: 0,
			MaxMsgSize:        0,
			Storage:           0,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func (ns *natsService) Publish(subjectName string, data []byte) error {
	i, err := ns.cfg.Js.Publish(subjectName, data)
	if err != nil {
		return err
	}
	log.Printf(":%v has been published\n", i.Stream)
	return nil
}
