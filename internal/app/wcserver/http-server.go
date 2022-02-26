package wcserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handleCrawl(w http.ResponseWriter, r *http.Request) {
	url, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	sitemap, err := crawlDomain(string(url))
	if err != nil {
		log.Printf("Error during crawling: %v", err)
		http.Error(w, "Error during crawling", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(sitemap)
	if err != nil {
		log.Printf("Error during marshalling result: %v", err)
		http.Error(w, "Error during marshalling", http.StatusInternalServerError)
		return
	}
}

func Serve() {
	http.HandleFunc("/", handleCrawl)
	fmt.Println("Listening on 0.0.0.0:80")
	http.ListenAndServe("", nil)
}
