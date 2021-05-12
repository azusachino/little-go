package main

import (
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
)

func main() {
	cfg := es.Config{
		Addresses: []string{"http://172.31.103.161:9200/"},
	}
	esClient, err := es.NewClient(cfg)

	if err != nil {
		log.Fatal(err)
	}
	esClient.Ping()

}
