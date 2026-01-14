package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	inputURL, _ = strings.CutSuffix(inputURL, "/")

	url, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("url parse err: %v", err)
	}

	return url.Host + url.Path, nil
}
