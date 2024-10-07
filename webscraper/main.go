package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	// "net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type ScrapedData struct {
	URL         string
	Title       string
	Description string
	Keywords    []string
	Links       []string
	Images      []string
	LastModified time.Time
}

type ScrapeStats struct {
	TotalPages     int
	TotalLinks     int
	TotalImages    int
	AverageKeywords float64
	DomainCount    map[string]int
	ExecutionTime  time.Duration
}

func main() {
	// Define command-line flags
	urlFlag := flag.String("url", "https://example.com", "The URL to scrape")
	outputFlag := flag.String("output", "scraped_data", "Output file name (without extension)")
	maxDepthFlag := flag.Int("depth", 2, "Maximum depth to crawl")
	formatFlag := flag.String("format", "csv", "Output format (csv or json)")
	concurrencyFlag := flag.Int("concurrency", 2, "Number of concurrent scrapers")
	flag.Parse()

	startTime := time.Now()

	// Create a new collector with configurations
	c := colly.NewCollector(
		colly.MaxDepth(*maxDepthFlag),
		colly.Async(true),
	)

	// Set up concurrency and delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: *concurrencyFlag,
		Delay:       1 * time.Second,
		RandomDelay: 500 * time.Millisecond,
	})

	var scrapedData []ScrapedData
	var stats ScrapeStats
	var mu sync.Mutex

	// Set up callbacks
	c.OnHTML("html", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()

		links := []string{}
		e.ForEach("a[href]", func(_ int, el *colly.HTMLElement) {
			link := el.Attr("href")
			if link != "" {
				links = append(links, link)
			}
		})

		images := []string{}
		e.ForEach("img[src]", func(_ int, el *colly.HTMLElement) {
			img := el.Attr("src")
			if img != "" {
				images = append(images, img)
			}
		})

		keywords := strings.Split(e.ChildAttr("meta[name=keywords]", "content"), ",")
		for i := range keywords {
			keywords[i] = strings.TrimSpace(keywords[i])
		}

		lastModified, _ := time.Parse(time.RFC1123, e.Response.Headers.Get("Last-Modified"))

		data := ScrapedData{
			URL:         e.Request.URL.String(),
			Title:       e.ChildText("title"),
			Description: e.ChildAttr("meta[name=description]", "content"),
			Keywords:    keywords,
			Links:       links,
			Images:      images,
			LastModified: lastModified,
		}
		scrapedData = append(scrapedData, data)

		// Update stats
		stats.TotalPages++
		stats.TotalLinks += len(links)
		stats.TotalImages += len(images)
		domain := e.Request.URL.Hostname()
		stats.DomainCount[domain]++
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %v", r.Request.URL, err)
	})

	// Initialize stats
	stats.DomainCount = make(map[string]int)

	// Start scraping
	err := c.Visit(*urlFlag)
	if err != nil {
		log.Fatal("Error visiting the site:", err)
	}

	// Wait for all goroutines to finish
	c.Wait()

	// Calculate final stats
	stats.ExecutionTime = time.Since(startTime)
	if len(scrapedData) > 0 {
		totalKeywords := 0
		for _, data := range scrapedData {
			totalKeywords += len(data.Keywords)
		}
		stats.AverageKeywords = float64(totalKeywords) / float64(len(scrapedData))
	}

	// Save data
	if *formatFlag == "json" {
		saveToJSON(scrapedData, stats, *outputFlag+".json")
	} else {
		saveToCSV(scrapedData, *outputFlag+".csv")
	}

	// Print summary
	printSummary(stats)
}

func saveToCSV(data []ScrapedData, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"URL", "Title", "Description", "Keywords", "Links", "Images", "Last Modified"})

	// Write data
	for _, item := range data {
		writer.Write([]string{
			item.URL,
			item.Title,
			item.Description,
			strings.Join(item.Keywords, "|"),
			strings.Join(item.Links, "|"),
			strings.Join(item.Images, "|"),
			item.LastModified.Format(time.RFC3339),
		})
	}
}

func saveToJSON(data []ScrapedData, stats ScrapeStats, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	output := struct {
		Data  []ScrapedData `json:"data"`
		Stats ScrapeStats   `json:"stats"`
	}{
		Data:  data,
		Stats: stats,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(output); err != nil {
		log.Fatal("Cannot encode to JSON", err)
	}
}

func printSummary(stats ScrapeStats) {
	fmt.Printf("\nScraping Summary:\n")
	fmt.Printf("Total Pages Scraped: %d\n", stats.TotalPages)
	fmt.Printf("Total Links Found: %d\n", stats.TotalLinks)
	fmt.Printf("Total Images Found: %d\n", stats.TotalImages)
	fmt.Printf("Average Keywords per Page: %.2f\n", stats.AverageKeywords)
	fmt.Printf("Domains Scraped:\n")
	for domain, count := range stats.DomainCount {
		fmt.Printf("  %s: %d pages\n", domain, count)
	}
	fmt.Printf("Execution Time: %v\n", stats.ExecutionTime)
}
