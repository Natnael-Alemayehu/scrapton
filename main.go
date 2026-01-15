package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(arg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Passed argument is not a valid url")
		os.Exit(1)
	}

	maxConcurrent := 5

	cfg := config{
		pages:              make(map[string]PageData),
		baseURL:            rawBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrent),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL.String())

	cfg.wg.Wait()

	fmt.Printf("\nCrawl Result:\n")
	for normalizedURL, page := range cfg.pages {
		fmt.Printf("   %s - %v\n", normalizedURL, len(page.OutgoingLinks))
	}

}
