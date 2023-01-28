package services

type pageInfo struct {
	StatusCode int
	Links      map[string]int
}

// Path: cmd\internals\services\colly_scrapper.go

func (p *pageInfo) GetStatusCode() int {
	return p.StatusCode
}

func (p *pageInfo) SetStatusCode(statusCode int) {

	p.StatusCode = statusCode
}

func (p *pageInfo) GetLinks() map[string]int {
	return p.Links
}

func (p *pageInfo) SetLinks(links map[string]int) {
	p.Links = links
}
