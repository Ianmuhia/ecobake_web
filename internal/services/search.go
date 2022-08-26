package services

import (
	"ecobake/cmd/config"
	"ecobake/internal/models"
	"encoding/json"

	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/meilisearch/meilisearch-go"
)

type searchService struct {
	cfg          *config.AppConfig
	searchClient *meilisearch.Client
}

type SearchService interface {
	IndexUser(bchan chan *models.User)
	IndexBusiness(bchan chan interface{})
	UpdateBusinessDoc(bchan chan interface{})
	UpdateUserDoc(bchan chan *models.User)
	DeleteBusinessDoc(delchan chan int32)
	SearchUser() (data interface{}, err error)
	DeleteAllDocs(doc string)
	CreateIndexes(fi models.FaindaIndexes, wg *sync.WaitGroup)
}

func NewSearchService(cfg *config.AppConfig, searchClient *meilisearch.Client) *searchService {
	return &searchService{cfg: cfg, searchClient: searchClient}
}

func (z *searchService) SearchUser() (data interface{}, err error) {
	//documents, err := z.searchClient.Index("users").DeleteAllDocuments()
	fmt.Println(runtime.NumCPU())

	documents, err := z.searchClient.Index("users").Search("a", &meilisearch.SearchRequest{
		Limit: 500,
		Sort: []string{
			"id:asc",
		},
		Filter: "_geoRadius(45.472735, 9.184019, 200000000000000000)",
	})
	if err != nil {
		return documents, err
	}
	return documents, nil
}

func (z *searchService) IndexUser(bchan chan *models.User) {
	data := <-bchan
	mj, _ := json.Marshal(data)
	documents, err := z.searchClient.Index("users").AddDocuments(mj, "id")
	if err != nil {
		z.cfg.ZincRcvChan <- err
	}
	z.cfg.ZincRcvChan <- documents
}

func (z *searchService) UpdateUserDoc(bchan chan *models.User) {
	data := <-bchan
	mj, _ := json.Marshal(data)
	documents, err := z.searchClient.Index("users").UpdateDocuments(mj, "id")
	if err != nil {
		z.cfg.ZincRcvChan <- err
	}
	z.cfg.ZincRcvChan <- documents
	z.cfg.Logger.Println("finished indexing user: %w", documents)
}

func (z *searchService) IndexBusiness(bchan chan interface{}) {
	data := <-bchan

	mj, _ := json.Marshal(data)
	documents, err := z.searchClient.Index("business").AddDocuments(mj, "id")
	if err != nil {
		z.cfg.ZincRcvChan <- err
	}
	z.cfg.ZincRcvChan <- documents
	z.cfg.Logger.Println("finished indexing business: %w", documents)
}

func (z *searchService) UpdateBusinessDoc(bchan chan interface{}) {
	data := <-bchan

	var f []interface{}
	f = append(f, data)
	g, _ := json.Marshal(f)

	response, err := z.searchClient.Index("business").UpdateDocuments(g, "id")
	if err != nil {
		z.cfg.ZincRcvChan <- err
	}
	z.cfg.ZincRcvChan <- response
	z.cfg.Logger.Println("finished indexing business: %w", response)
}

func (z *searchService) DeleteBusinessDoc(delchan chan int32) {
	id := <-delchan

	documents, err := z.searchClient.Index("business").DeleteDocument(string(id))
	if err != nil {
		z.cfg.ZincRcvChan <- err
	}
	z.cfg.ZincRcvChan <- documents

	z.cfg.Logger.Println("finished deleting business _doc: %w", documents)

}

func (z *searchService) DeleteAllDocs(doc string) {

	documents, err := z.searchClient.Index(doc).DeleteAllDocuments()
	if err != nil {
		z.cfg.ZincRcvChan <- err
	}
	z.cfg.ZincRcvChan <- documents

	z.cfg.Logger.Println("finished deleting all _doc: %w", documents)

}

func (z *searchService) CreateIndexes(fi models.FaindaIndexes, wg *sync.WaitGroup) {
	taskInfo, err := z.searchClient.CreateIndex(
		&meilisearch.IndexConfig{PrimaryKey: fi.PrimaryKey, Uid: fi.Uid})
	if err != nil {
		log.Println(err)
	}
	g, err := z.searchClient.Index(fi.Uid).UpdateSearchableAttributes(fi.Searchableattributes)
	if err != nil {
		log.Println(err)
	}
	h, err := z.searchClient.Index(fi.Uid).UpdateFilterableAttributes(fi.Filterableattributes)
	if err != nil {
		log.Println(err)
	}
	d, err := z.searchClient.Index(fi.Uid).UpdateSortableAttributes(fi.Sortableattributes)
	if err != nil {
		log.Println(err)
	}
	z.cfg.Logger.Println(taskInfo.Details.IndexedDocuments)
	z.cfg.Logger.Println(g)
	z.cfg.Logger.Println(h)
	z.cfg.Logger.Println(d)
	z.cfg.Logger.Println(taskInfo.Error.Message)
	wg.Done()
}
