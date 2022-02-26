package wcserver

import (
	"bytes"
	"golang.org/x/net/html"
	"net/url"
)

func extractLinks(htmlPage []byte) ([]*url.URL, error) {
	doc, err := html.Parse(bytes.NewReader(htmlPage))
	if err != nil {
		return nil, err
	}
	var f func(*html.Node) ([]*url.URL, error)
	f = func(n *html.Node) ([]*url.URL, error) {
		var links []*url.URL
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := url.Parse(a.Val)
					if err != nil {
						return nil, err
					}
					links = append(links, link)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			childLinks, err := f(c)
			if err != nil {
				return nil, err
			}
			links = append(links, childLinks...)
		}
		return links, nil
	}
	return f(doc)
}
