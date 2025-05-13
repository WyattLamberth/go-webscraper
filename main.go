package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func scrapeURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	html, err := doc.Html()
	if err != nil {
		return "", err
	}

	return html, nil
}

func parse_args() []string {
	if len(os.Args) < 1 { // need at least one url
		fmt.Println("Usage: go run main.go <url1> <url2> ...")
		os.Exit(1)
	}
	return os.Args[1:]
}

func main() {
	urls := parse_args()

	type result struct {
		url  string
		html string
		err  error
	}

	results := make(chan result)

	for _, url := range urls {
		go func(u string) {
			fmt.Println("Scraping:", u)
			html, err := scrapeURL(u)
			results <- result{url: u, html: html, err: err}
		}(url)
	}

	for range urls {
		r := <-results
		if r.err != nil {
			fmt.Printf("Error scraping %s: %v\n", r.url, r.err)
			continue
		}
		fmt.Printf("✅ %s\n%s\n\n", r.url, r.html[:300]) // Trim to 300 chars for readability
	}
}
