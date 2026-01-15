package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !strings.Contains(rawCurrentURL, rawBaseURL) {
		return
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("rawBaseURL not parsing: %v \n", err)
		return
	}

	url, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url: %v \n", err)
		return
	}

	_, ok := pages[url]
	if ok {
		pages[url]++
		return
	}
	pages[url] = 1

	html, err := getHTML("https://" + url)
	if err != nil {
		fmt.Printf("Error getHTML: %v \n", err)
		return
	}

	fmt.Printf("Crawling URL: %v \n", url)

	urls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("Error getting urls: %v \n", err)
	}

	for _, v := range urls {
		crawlPage(rawBaseURL, v, pages)
	}

}
