package main

import (
	"fmt"
	"net/url"
	"os"
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

	pages := make(map[string]int)

	crawlPage(rawBaseURL.String(), rawBaseURL.String(), pages)

	fmt.Printf("\nCrawl Result:\n")
	for normalizedURL, count := range pages {
		fmt.Printf("   %d - %s\n", count, normalizedURL)
	}

}
