package connections

import (
	es7 "github.com/elastic/go-elasticsearch/v7"
)

func NewElasticSearchClient(addresses []string) (*es7.Client, error) {
	return es7.NewClient(es7.Config{ Addresses: addresses })
}