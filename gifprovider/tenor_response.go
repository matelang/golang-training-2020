package gifprovider

type TenorGifObject struct {
	WebUrl  string `json:"weburl"`
	Results []struct {
		HasCaption bool    `json:"hascaption"`
		Created    float64 `json:"created"`
		URL        string  `json:"url"`
		Tags      []interface{} `json:"tags"`
		Shares    int           `json:"shares"`
		ItemUrl   string        `json:"itemurl"`
		Composite interface{}   `json:"composite"`
		HasAudio  bool          `json:"hasaudio"`
		Title     string        `json:"title"`
		ID        string        `json:"id"`
	} `json:"results"`
	Next string `json:"next"`
}
