package engine

import (
	"errors"
	"github.com/matelang/gifseek/gifprovider"
	log "github.com/sirupsen/logrus"
)

type SequentialSearchEngine struct {
	GifProviders []gifprovider.GifProvider
}

func (s SequentialSearchEngine) Search(query string) (*SearchResults, error) {
	finalSearchResults := &SearchResults{}

	for _, p := range s.GifProviders {
		result, err := p.GetGifs(query)

		if err != nil {
			log.Warnf("Could not get GIFs from provider %s for query=%s", p, query)
			continue
		}

		finalSearchResults.Urls = append(finalSearchResults.Urls, result.Urls...)
	}

	if len(finalSearchResults.Urls) == 0 {
		return nil, errors.New("could not retrieve any GIFs from any provider")
	}

	return finalSearchResults, nil
}
