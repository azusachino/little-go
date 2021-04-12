package main

import "github.com/gocolly/colly/v2"

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		element.Request.Visit(element.Attr("href"))
	})
}
