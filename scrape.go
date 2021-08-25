package main

import ( 
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	var count int

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error){
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(".s-item__info",func(listing *colly.HTMLElement){
		listing_title := listing.ChildText(".s-item__title")
		listing_price := listing.ChildText(".s-item__price")
		condition := listing.ChildText(".SECONDARY_INFO")
		Shipping := listing.ChildText(".s-item__shipping")
		Pickup := listing.ChildText(".s-item__dynamic")
		if len(listing_title) > 1{
			fmt.Println("\nTitle: " + string(listing_title) + "\nPrice: " + listing_price + "\nCondition: " + condition)
			count++
		}
		if len(Shipping) > 0 {
			fmt.Println(Shipping)
		} else if len(Pickup) > 0 {
			fmt.Println(Pickup)
		}
	})

	c.Visit("https://www.ebay.com/sch/i.html?_from=R40&_nkw=se%20bikes%20vans&_sacat=0&LH_TitleDesc=0&rt=nc&_udlo=500")

	fmt.Println("\nListings found: ")
	fmt.Print(count)

}
