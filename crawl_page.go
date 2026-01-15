package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("rawBaseURL not parsing: %v \n", err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url: %v \n", err)
		return
	}

	// increment if visited
	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getHTML: %v \n", err)
		return
	}

	Nexturls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("Error getting urls: %v \n", err)
	}

	for _, nextURL := range Nexturls {
		crawlPage(rawBaseURL, nextURL, pages)
	}

}
