package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", errors.New("error getting response. Getting status code > 400")
	}

	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", errors.New("not getting text/html")
	}

	pageHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(pageHTML), nil
}
