package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(arg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	if os.Args[1] == "help" {
		fmt.Println("./crawller <URL> <maxConcurrent> <maxPages>")
		fmt.Println("Default maxConcurrent: 5, maxPages: 10")
		os.Exit(0)
	}

	maxConcurrent := 5
	maxPages := 10

	rawBaseURL, err := url.Parse(arg[0])
	if err != nil {
		fmt.Println("Passed argument is not a valid url")
		os.Exit(1)
	}

	if len(arg) > 1 {
		maxConcurrentArg := arg[1]
		maxConcurrent, err = strconv.Atoi(maxConcurrentArg)
		if err != nil {
			fmt.Println("max concurrency must be valid integer")
		}

		maxPagesArgs := arg[2]
		maxPages, err = strconv.Atoi(maxPagesArgs)
		if err != nil {
			fmt.Println("max concurrency must be valid integer")
		}

	}

	cfg := config{
		pages:              make(map[string]PageData),
		baseURL:            rawBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrent),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL.String())

	cfg.wg.Wait()

	timestamp := time.Now().Format("2006-01-02_15-04-05")

	filename := fmt.Sprintf("./exports/report_csv_%v.csv", timestamp)

	if err := writeCSVReport(cfg.pages, filename); err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

}
