package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println("No data to write to CSV")
		return nil
	}

	fname, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Create file error: %v", err)
	}
	defer fname.Close()

	csvwriter := csv.NewWriter(fname)
	defer csvwriter.Flush()

	header := []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}
	if err := csvwriter.Write(header); err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, normalizedURL := range keys {
		p := pages[normalizedURL]
		outgoing := strings.Join(p.OutgoingLinks, ";")
		images := strings.Join(p.ImageURLs, ";")
		row := []string{
			p.URL,
			p.H1,
			p.FirstParagraph,
			outgoing,
			images,
		}
		if err := csvwriter.Write(row); err != nil {
			return fmt.Errorf("write row for %s: %w", p.URL, err)
		}
	}

	fmt.Printf("\n\nReport written to %s\n", filename)

	return nil
}
