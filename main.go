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

	url, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Passed argument is not a valid url")
		os.Exit(1)
	}

	pages := map[string]int{}

	crawlPage(url.String(), url.String(), pages)

	fmt.Println("Done crawling!")
	fmt.Println("crawl result: ")

	for k, v := range pages {
		fmt.Printf("\t- url: %v - %v \n", k, v)
	}

}
