package wcclient

import (
	"encoding/json"
	"fmt"
	"github.com/ebr41nd/web-crawler/internal/pkg/crawl"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var wcServerHost string

func init() {
	wcServerHost = os.Getenv("WC_SERVER_HOST")
	if wcServerHost == "" {
		fmt.Println("Defaulting WC_SERVER_HOST to http://localhost")
		wcServerHost = "http://localhost"
	}
}

func Execute() {
	if len(os.Args) < 2 {
		log.Fatal("wcclient needs an url as only argument")
		os.Exit(1)
	}
	page := os.Args[1]
	fmt.Printf("Call the web-crawler server at %s\n", wcServerHost)
	r, err := http.Post(wcServerHost, "application/x-www-form-urlencoded", strings.NewReader(page))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var crawlResult crawl.CrawlResult
	json.Unmarshal(result, &crawlResult)
	printResult(&crawlResult, 0)
}

func printResult(result *crawl.CrawlResult, ident int) {
	if ident == 0 {
		fmt.Printf("%s\n", result.Url)
	} else {
		fmt.Printf("%s- %s\n", strings.Repeat(" ", ident), result.Url)
	}
	for _, link := range result.Links {
		printResult(link, ident+4)
	}
}
