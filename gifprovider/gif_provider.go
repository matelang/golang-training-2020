package gifprovider

import "fmt"

type ProviderSearchResults struct {
	ProviderId string
	Urls       []string
}

func (p ProviderSearchResults) String() string {
	return fmt.Sprintf("Provider=%s Urls=%v", p.ProviderId, p.Urls)
}

type GifProvider interface {
	GetGifs(query string) (*ProviderSearchResults, error)
}
