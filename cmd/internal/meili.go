package internal

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
)

type MeiliSearch struct {
	MeiliClient *meilisearch.Client
}

func NewMeiliSearch(meiliClient *meilisearch.Client) *MeiliSearch {
	return &MeiliSearch{MeiliClient: meiliClient}
}

func (mc *MeiliSearch) About() {
	health, err := mc.MeiliClient.Health()
	if err != nil {
		return
	}
	log.Println(health)

}

func GetMeiliConn() *meilisearch.Client {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})

	return client
}
