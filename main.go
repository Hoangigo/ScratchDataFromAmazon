package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Item struct {
	Name       string
	Price      string
	NumReviews string
	ImageURL   string
}

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Visiting %s\n", e.Error())

	})
	var products []Item
	c.OnHTML("div.zg-carousel-general-faceout", func(e *colly.HTMLElement) {

		product := Item{
			Name:       e.ChildText("a.a-link-normal > span > div.p13n-sc-truncate-desktop-type2"),
			Price:      e.ChildText("span.a-size-base.a-color-price span._cDEzb_p13n-sc-price_3mJ9Z"),
			NumReviews: e.ChildText("div.a-icon-row span.a-size-small"),
			ImageURL:   e.ChildAttr("img.a-dynamic-image", "src"),
		}
		fmt.Println(product)
		products = append(products, product)
	})
	err := c.Visit("https://www.amazon.de/gp/bestsellers/?ref_=nav_cs_bestsellers")

	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = ioutil.WriteFile("products.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data saved to products.json")
}
