package gifprovider

const GiphyProviderId = "Giphy"

type GiphyGifProvider struct {
}

func (g GiphyGifProvider) GetGifs(query string) (*ProviderSearchResults, error) {

	return &ProviderSearchResults{
		ProviderId: GiphyProviderId,
		Urls:       []string{"https://giphy.com/cat"},
	}, nil
}
