package api

import (
	"log"

	"github.com/gocolly/colly"
)

func ScrapeRedfinAPI() {
	// Create a new collector
	c := colly.NewCollector()

	// Use the collector as a proxy
	c.SetProxy("http://localhost:8080")

	// Scrape the Redfin API
	c.OnResponse(func(r *colly.Response) {
		log.Println("Response received", string(r.Body))
	})

	err := c.Visit("https://api.redfin.com/api/apiresults/search")
	if err != nil {
		log.Fatalf("Visit failed: %v", err)
	}
}
