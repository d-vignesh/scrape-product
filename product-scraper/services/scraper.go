package services

import (
	"strings"
	"strconv"
	"reflect"
	"encoding/json"

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

	// checking for the product title
	c.OnHTML(`span[id=productTitle]`, func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		s.logger.Info("found title", title)
		prod.Name = title
	})

	// checking for the product image url
	c.OnHTML(`div[id=imgTagWrapperId]`, func(e *colly.HTMLElement) {
		imgs := e.ChildAttrs("img", "data-a-dynamic-image")
		// imgs is a list conainting one string which is a json object in the form url: pixel of all the images of the product
		if len(imgs) > 0 {
			imgStr := imgs[0]
			imgData := make(map[string][]int)
			if err := json.Unmarshal([]byte(imgStr), &imgData); err != nil {
				s.logger.Error("unable to unmarshal images list", "error", err)
			} else {
				prod.ImageURL = reflect.ValueOf(imgData).MapKeys()[0].String()
				s.logger.Info("found image url", prod.ImageURL)
			}
		} else {
			s.logger.Error("could not find image url")
		}
	})

	// checking for the product description
	c.OnHTML(`div[id=productDescription]`, func(e *colly.HTMLElement) {
		description := e.ChildText("p")
		s.logger.Info("found description", description)
		prod.Description = description
	})

	// checking for the product description provided as 'About the product'
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

	// checking for product reviews count
	c.OnHTML(`span[id=acrCustomerReviewText]`, func(e *colly.HTMLElement) {
		if prod.TotalReviews == 0 {
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

	// check for product price, the price was provided under different tags for different tags, 
	// we are checking for all the possible tags.
	c.OnHTML(`span[id=priceblock_saleprice]`, func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		s.logger.Info("found price", price)
		prod.Price = price	
	})

	c.OnHTML(`span[id=priceblock_ourprice]`, func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		s.logger.Info("found price", price)
		prod.Price = price
	})

	c.OnHTML(`span[id=priceblock_dealprice]`, func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		s.logger.Info("found price", price)
		prod.Price = price
	})

	c.Visit(url)

	return &prod
}