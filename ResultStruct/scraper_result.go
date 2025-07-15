package ResultStruct

type ScraperResult struct {
	Result []Result `json:"results"`
}

type Result struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	PubDate string `json:"pubDate"`
}
