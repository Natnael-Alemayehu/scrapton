# Scrapton

Scrapton is a small, concurrent web crawler written in Go that visits pages on a single host, extracts basic page information (H1, first paragraph, outbound links, and image URLs), and produces a CSV report for analysis.

**Table of Contents**
- **Project:** Overview and goals
- **Install:** Requirements and setup
- **Usage:** How to run the crawler and available arguments
- **Output:** Report format and example
- **Project structure:** Key source files
- **Testing:** Running the existing tests
- **Contributing & License**

**Project**

Scrapton is designed as a lightweight site crawler to collect lightweight metadata from HTML pages. It crawls only pages that share the same hostname as the provided start URL, supports configurable concurrency and page limits, and writes results to a CSV file under the `exports/` directory.

**Install**

Prerequisites:
- Go toolchain (Go 1.20+ is recommended; the repository declares a more recent Go version in `go.mod`).

Quick setup:

```bash
git clone https://github.com/<your-username>/scrapton.git
cd scrapton
go mod tidy
```

Build the binary (optional):

```bash
go build -o crawller
```

Or run directly with `go run`:

```bash
go run . -- https://example.com 5 10
```

**Usage**

The program accepts positional arguments:

- `<URL>` (required): The start URL to crawl. Scrapton will only crawl pages under this URL's hostname.
- `<maxConcurrent>` (optional): Maximum number of concurrent requests (default: 5).
- `<maxPages>` (optional): Maximum number of pages to collect (default: 10).

Example:

```bash
./crawller https://example.com 5 50
```

Notes:
- The built-in help shows the usage string: `./crawller <URL> <maxConcurrent> <maxPages>`
- If you omit `maxConcurrent` and `maxPages`, Scrapton uses sensible defaults (5 and 10 respectively).

**Output**

When the crawl completes, Scrapton writes a CSV report to `./exports/report_csv_<unix_timestamp>.csv` containing the following columns:

- `page_url`: Page URL as seen by the crawler.
- `h1`: The first H1 text on the page if present.
- `first_paragraph`: The first paragraph text (prefers content inside `<main>` when available).
- `outgoing_link_urls`: Semicolon-separated list of resolved outbound links found on that page.
- `image_urls`: Semicolon-separated list of resolved image `src` values found on that page.

The CSV writer sorts pages by their normalized URL to produce a deterministic ordering.

**Project structure (key files)**

- `main.go` — CLI entry point; parses args, initializes crawler config, and writes the CSV report.
- `crawl_page.go` — Core concurrent crawler and page scheduling logic.
- `pagedata.go` — PageData struct and the `extractPageData` function that coordinates data extraction.
- `get_data_from_html.go` — HTML parsing helpers (H1, first paragraph, links, images) using `goquery`.
- `getHTML.go` — HTTP client and page fetch logic with basic content-type and status checks.
- `csv_report.go` — CSV report generation logic.
- `exports/` — Target directory for generated CSV reports.

**Dependencies**

The project uses `github.com/PuerkitoBio/goquery` for HTML parsing. Dependencies are declared in `go.mod`; fetch them with:

```bash
go mod tidy
```

**Testing**

Run the test suite with:

```bash
go test ./...
```

There are unit tests for URL normalization and HTML extraction helpers in the repository.

**Contributing**

Contributions are welcome. For code changes, please:

1. Fork the repository.
2. Create a feature branch for your change.
3. Open a pull request describing the change and rationale.

If you plan to add features, consider:
- Adding command-line flags for named arguments (instead of positional args).
- Respecting `robots.txt` before crawling.
- Adding rate limiting and retry/backoff policies.
- Adding more robust error handling and logging.

**License & Contact**

Specify your preferred license (for example, MIT) for this repository. For questions or feedback, open an issue in this repository.

---

This README was generated to provide a concise, professional overview of the project and how to use it. If you would like, I can also add a CONTRIBUTING.md, LICENSE, or example output CSV to the `exports/` folder.
