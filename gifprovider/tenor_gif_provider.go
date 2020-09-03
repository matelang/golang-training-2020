package gifprovider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const TenorProviderId = "Tenor"
const TenorApiKey = "" //TODO insert your api key
const Limit = 10

type TenorGifProvider struct {
}

func (t TenorGifProvider) GetGifs(query string) (*ProviderSearchResults, error) {
	getString := fmt.Sprintf("https://api.tenor.com/v1/search?key=%s&q=%s&limit=%d",
		TenorApiKey, url.QueryEscape(query), Limit)

	resp, err := http.Get(getString)
	if err != nil {
		return nil, err
	}

	tenorGifObject := &TenorGifObject{}
	if err := json.NewDecoder(resp.Body).Decode(tenorGifObject); err != nil {
		return nil, err
	}

	resultUrls := make([]string, len(tenorGifObject.Results))
	for i, r := range tenorGifObject.Results {
		resultUrls[i] = r.URL
	}

	return &ProviderSearchResults{
		ProviderId: TenorProviderId,
		Urls:       resultUrls,
	}, nil
}
