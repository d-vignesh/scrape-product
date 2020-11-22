package services

import (
	"strings"
	"strconv"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-scraper/data"
)

type Scraper struct {
	logger hclog.Logger
}

func NewScraper(logger hclog.Logger) *Scraper {
	return &Scraper{logger}
}

func (s *Scraper) ScrapeURL(url string) *data.Product {
	
	prod := data.Product{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		s.logger.Info("visiting", r.URL)
	})

	c.OnHTML(`span[id=productTitle]`, func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		s.logger.Info("found title", title)
		prod.Name = title
	})

	c.OnHTML(`img[id=landingImage]`, func(e *colly.HTMLElement) {
		imgURL := e.Attr("src")
		s.logger.Info("found img url", imgURL)
		s.logger.Info("url class", e.Attr("alt"))
		prod.ImageURL = imgURL
	})

	c.OnHTML(`div[id=productDescription]`, func(e *colly.HTMLElement) {
		description := e.ChildText("p")
		s.logger.Info("found description", description)
		prod.Description = description
	})

	c.OnHTML(`div[id=feature-bullets]`, func(e *colly.HTMLElement) {
		if prod.Description == "" {
			var sb strings.Builder
			e.ForEach("ul > li > span", func(_ int, elem *colly.HTMLElement) {
				sb.WriteString(strings.TrimSpace(elem.Text))
			})

			description := sb.String()
			s.logger.Info("found description", description)
			prod.Description = description
		}
	})

	c.OnHTML(`span[id=acrCustomerReviewText]`, func(e *colly.HTMLElement) {
		if prod.Reviews == 0 {
			ratingString := strings.TrimSpace(e.Text)
			s.logger.Info("found ratings", ratingString)
			ratings := strings.Split(ratingString, " ")
			if len(ratings) > 1 {
				ratingCount, err := strconv.Atoi(strings.Replace(ratings[0], ",", "", -1))
				if err != nil {
					s.logger.Error("could not parse rating count")
				} else {
					prod.TotalReviews = ratingCount
				}
			} else {
				s.logger.Error("unable to parse ratings string", ratingString)
			}
		}
	})

	c.OnHTML(`span[id=priceblock_saleprice]`, func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		s.logger.Info("found price", priceString)
		prod.Price = price	
	})

	c.OnHTML(`span[id=priceblock_ourprice]`, func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		s.logger.Info("found price", priceString)
		prod.Price = price
	})

	c.OnHTML(`span[id=priceblock_dealprice]`, func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		s.logger.Info("found price", priceString)
		prod.Price = price
	})

	c.Visit(url)

	return &prod
}