package wcserver

import (
	"bytes"
	"encoding/json"
	"github.com/ebr41nd/web-crawler/internal/pkg/crawl"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestCrawlDomain(t *testing.T) {
	client = FakeClient{}
	got, err := crawlDomain("http://example.com")
	fData, _ := readTestFile("example.com.json")
	want := &crawl.CrawlResult{}
	json.Unmarshal(fData, &want)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(got, want) {
		t.Log("The sitemap is not the one expected want :")
		printCrawlResult(got, t)
		printCrawlResult(want, t)
	}

}

func TestIsSubLink(t *testing.T) {
	url1, _ := url.Parse("http://example.com/")
	url2, _ := url.Parse("http://example.com/tokeep/")
	got := isSubLink(url1, url2)

	if !got {
		t.Errorf("%v is not a sub link of %v", url2, url1)
	}
}

func TestIsSubLinkNotSubPath(t *testing.T) {
	url1, _ := url.Parse("http://example.com/token")
	url2, _ := url.Parse("http://example.com/tokeep/")
	got := isSubLink(url1, url2)

	if got {
		t.Errorf("%v is a sub link of %v", url2, url1)
	}
}

func TestIsSubLinkNotSameDomain(t *testing.T) {
	url1, _ := url.Parse("http://example.com/token")
	url2, _ := url.Parse("http://examples.com/token/test")
	got := isSubLink(url1, url2)

	if got {
		t.Errorf("%v is a sub link of %v", url2, url1)
	}
}

func TestFilterAnchors(t *testing.T) {
	var links []*url.URL
	url1, _ := url.Parse("http://example.com/#todelete")
	url2, _ := url.Parse("http://example.com/tokeep/")
	links = append(links, url1, url2)
	got := filterAnchors(links)

	if contains(got, "#todelete") {
		t.Errorf("#todelete is still in links: %v", links)
	}
}

type FakeClient struct{}

func (f FakeClient) Get(u string) (resp *http.Response, err error) {
	var data []byte
	parsedUrl, _ := url.Parse(u)
	filename := parsedUrl.Host
	if len(parsedUrl.Path) > 0 {
		filename += strings.ReplaceAll(parsedUrl.Path, "/", "-")
		filename = strings.TrimSuffix(filename, "-")
	}
	filename += ".html"
	data, err = readTestFile(filename)
	if err != nil {
		return nil, err
	}
	return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(data))}, nil
}

func printCrawlResult(c *crawl.CrawlResult, t *testing.T) {
	t.Log(c.Url)
	for _, l := range c.Links {
		printCrawlResult(l, t)
	}
}
