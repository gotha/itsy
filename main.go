package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	//"github.com/elastic/go-elasticsearch/v7"
)

var startURL string
var cookie string

func main() {
	flag.StringVar(&startURL, "startURL", "", "the first page where we should start scraping")
	flag.StringVar(&cookie, "cookie", "", "if provided will be sent to host")
	flag.Parse()

	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}

	req, err := http.NewRequest("GET", startURL, nil)
	if err != nil {
		log.Panicf("could not create http request: %s", err)
	}
	req.Header.Set("cookie", cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("could not load start URL: %s", err)
	}

	if resp.StatusCode != 200 {
		log.Panicf("unexpected http status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	getAllLinksFromPage(resp.Body)

	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Panicf("could not read body on start page: %s", err)
	//}

	//fmt.Printf("%+v\n", string(body))
}

func getAllLinksFromPage(body io.ReadCloser) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		fmt.Printf("%+v\n", href)
	})
}
