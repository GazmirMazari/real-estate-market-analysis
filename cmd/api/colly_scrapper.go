package redfin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Property represents a property listed on Redfin
type Property struct {
	Address string
	City    string
	State   string
	Zip     string
	Price   string
}

func Scrape(url string) ([]Property, error) {
	var properties []Property

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	doc.Find(".HomeCardContainer").Each(func(i int, s *goquery.Selection) {
		address := strings.TrimSpace(s.Find(".Address").Text())
		price := strings.TrimSpace(s.Find(".HomePrice").Text())

		var city, state, zip string
		location := strings.Split(strings.TrimSpace(s.Find(".Address").Text()), ", ")
		if len(location) == 3 {
			city = location[0]
			state = location[1]
			zip = location[2]
		}

		property := Property{
			Address: address,
			City:    city,
			State:   state,
			Zip:     zip,
			Price:   price,
		}
		properties = append(properties, property)
	})

	return properties, nil
}
