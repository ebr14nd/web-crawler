package crawl

type CrawlResult struct {
	Url   string         `json:"url"`
	Links []*CrawlResult `json:"links,omitempty"`
}
