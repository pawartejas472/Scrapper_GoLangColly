package scrapper

import (
	"fmt"
	"log"
	"strconv"
	"github.com/gocolly/colly"
)

type Product struct {
	Title        string `json:"id"`
	ImageURL     string `josn:"imageurl"`
	Dscription   string `json:"description"`
	Price        int    `json:"price"`
	NoOfReviews  string `json:"noofreviews`
}

func main() {


}

func scrape(url string ) {
	allProducts := make([]Product, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("https://www.amazon.com", "www.amazon.in"),
	)

	collector.OnHTML("div.s-main-slot s-result-list s-search-results sg-row", func(element *colly.HTMLElement) {
		element.ForEach("div.a-section.a-spacing-medium", func(_ int, element *colly.HTMLElement) {
			var productTitle string
			productTitle = element.ChildText("span.a-size-medium a-color-base a-text-normal")

			var productImageUrl string
			productImageUrl = element.ChildAttr(`img`,"src")

			var productDesc string
			productDesc = element.ChildText("div.a-row a-size-base a-color-secondary s-align-children-center")
			
			var productPrice int
			productPrice = element.ChildText("span.a-offscreen")
			productPrice = strconv.Atoi(productPrice)
			
			var productNoOfReviews string
			productNoOfReviews = element.ChildText("a.a-link-normal")
			
			oneProduct := Product{
				Title       : productTitle 
				ImageURL    : productImageUrl 
				Dscription  : productDesc
				Price       : productPrice
				NoOfReviews : productNoOfReviews

			}
			allProducts = append(allProducts, oneProduct)
		}

	}
	
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	}
	collector.Visit("https://www.amazon.in/s?k=jbl+earphone%E0%A4%82&crid=2PRT6R0JU03VS&sprefix=jbl+ear%2Caps%2C549&ref=nb_sb_ss_ts-doa-p_1_7")
}
