package wcserver

import (
	"fmt"
	"github.com/ebr41nd/web-crawler/internal/pkg/crawl"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

var client httpClient

func init() {
	client = http.DefaultClient
}

type domainToCrawl struct {
	urlPart *url.URL
	fullUrl *url.URL
}

func crawlDomain(domain string) (*crawl.CrawlResult, error) {
	baseUrl, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	toCrawl := domainToCrawl{baseUrl, baseUrl}
	result := &crawl.CrawlResult{}
	var wg sync.WaitGroup
	wg.Add(1)
	internalCrawl(&toCrawl, result, &wg)
	wg.Wait()
	return result, nil
}

func internalCrawl(toCrawl *domainToCrawl, result *crawl.CrawlResult, wg *sync.WaitGroup) {
	fmt.Printf("Crawling: %v\n", toCrawl.fullUrl)
	result.Url = toCrawl.urlPart.String()
	links, err := fetchLinks(toCrawl.fullUrl)
	if err != nil {
		log.Println(err)
	}
	links = filterAnchors(links)
	var swg sync.WaitGroup
	for _, link := range links {
		subUrl := buildUrl(toCrawl.fullUrl, link)
		if !isSubLink(toCrawl.fullUrl, subUrl) {
			continue
		}
		subResult := &crawl.CrawlResult{}
		result.Links = append(result.Links, subResult)
		swg.Add(1)
		go internalCrawl(&domainToCrawl{link, subUrl}, subResult, &swg)
	}
	swg.Wait()
	wg.Done()
}

func isSubLink(u *url.URL, su *url.URL) bool {
	return u.Host == su.Host && su.Path != u.Path && strings.Contains(su.Path, u.Path)
}

func buildUrl(u *url.URL, su *url.URL) *url.URL {
	return u.ResolveReference(su)
}

func filterAnchors(links []*url.URL) []*url.URL {
	var filteredLinks []*url.URL
	for _, link := range links {
		if !strings.HasPrefix(link.Path, "#") {
			filteredLinks = append(filteredLinks, link)
		}
	}
	return filteredLinks
}
