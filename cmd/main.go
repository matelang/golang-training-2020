package main

import (
	"flag"
	"github.com/matelang/gifseek/engine"
	"github.com/matelang/gifseek/gifprovider"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	tenorGifProvider := &gifprovider.TenorGifProvider{}
	giphyGifProvider := &gifprovider.GiphyGifProvider{}

	searchEngine := &engine.SequentialSearchEngine{
		GifProviders: []gifprovider.GifProvider{tenorGifProvider, giphyGifProvider},
	}

	query := flag.String("query", "", "Query string")
	flag.Parse()

	results, err := searchEngine.Search(*query)
	if err != nil {
		log.Fatalf("Search engine failed %s", err)
	}

	for i, u := range results.Urls {
		log.Infof("Url %d -> %s", i, u)
	}

}
