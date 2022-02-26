package wcserver

import (
	"io/ioutil"
	"net/url"
)

func fetchLinks(u *url.URL) ([]*url.URL, error) {
	page, err := fetchPage(u)
	if err != nil {
		return nil, err
	}
	links, err := extractLinks(page)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func fetchPage(u *url.URL) ([]byte, error) {
	r, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}
	page, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return page, nil
}
