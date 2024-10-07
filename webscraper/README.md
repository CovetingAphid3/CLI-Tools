# Advanced Web Scraper

This is an advanced web scraper built in Go, designed to efficiently crawl websites, extract valuable information, and provide detailed analytics. It's highly configurable and outputs data in both CSV and JSON formats.

## Features

- Concurrent scraping with configurable number of workers
- Depth-limited crawling
- Extracts comprehensive page data:
  - URL
  - Title
  - Description (from meta tags)
  - Keywords (from meta tags)
  - All links on the page
  - All images on the page
  - Last modified date (from HTTP headers)
- Respects robots.txt and implements polite crawling with configurable delays
- Outputs data in CSV or JSON format
- Provides detailed scraping statistics and analytics
- Command-line interface for easy configuration

## Requirements

- Go 1.15 or higher
- github.com/gocolly/colly/v2 package

## Installation

1. Clone this repository:
   ```
   cd advanced-web-scraper
   ```

2. Install the required dependencies:
   ```
   go mod tidy
   ```

## Usage

Run the scraper using the following command:

```
go run main.go [flags]
```

### Flags

- `-url`: The URL to start scraping from (default: "https://example.com")
- `-output`: Output file name without extension (default: "scraped_data")
- `-format`: Output format, either "csv" or "json" (default: "csv")
- `-depth`: Maximum depth to crawl (default: 2)
- `-concurrency`: Number of concurrent scrapers (default: 2)

### Examples

1. Scrape a website with default settings:
   ```
   go run main.go -url=https://example.com
   ```

2. Scrape with custom depth and output format:
   ```
   go run main.go -url=https://example.com -depth=3 -format=json -output=results
   ```

3. Scrape with increased concurrency:
   ```
   go run main.go -url=https://example.com -concurrency=5
   ```

## Output

### CSV Format

The CSV output includes the following columns:
- URL
- Title
- Description
- Keywords (pipe-separated)
- Links (pipe-separated)
- Images (pipe-separated)
- Last Modified

### JSON Format

The JSON output includes two main sections:
1. `data`: An array of scraped pages, each containing all extracted information.
2. `stats`: Overall scraping statistics including total pages, links, images, average keywords per page, domain distribution, and execution time.

## Scraping Summary

After completing the scrape, the program will print a summary to the console, including:
- Total pages scraped
- Total links found
- Total images found
- Average keywords per page
- List of scraped domains with page counts
- Total execution time


