package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

// A Devtools driver for web automation and scraping

func main() {
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")

	page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#search-form > fieldset > button").MustClick()

	el := page.MustElement("#mw-content-text > div.mw-parser-output > table.infobox > tbody > tr:nth-child(1) > td > a > img")
	_ = utils.OutputFile("b.png", el.MustResource())
}
