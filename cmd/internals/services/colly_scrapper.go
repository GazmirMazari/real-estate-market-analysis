package services

import (
	"github.com/gocolly/colly"
)

type pageInfo struct {
	StatusCode int
	Links      map[string]int
	URL        string
}

func (p *pageInfo) scrappeZillow(url string, data chan pageInfo) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.zillow.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		p.Links[link]++
	})

	c.OnResponse(func(r *colly.Response) {
		p.StatusCode = r.StatusCode
		p.URL = url
	})

	c.Visit(url)
	data <- *p
}
