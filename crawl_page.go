package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url: %v \n", err)
		return
	}

	cfg.mu.Lock()
	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.mu.Unlock()
		return
	}

	cfg.pages[normalizedURL] = PageData{}
	cfg.mu.Unlock()

	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getHTML: %v \n", err)
		return
	}

	pd := extractPageData(html, rawCurrentURL)

	cfg.mu.Lock()
	cfg.pages[normalizedURL] = pd
	cfg.mu.Unlock()

	for _, link := range pd.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

}
