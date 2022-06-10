package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Data struct {
	title       string
	description string
}

func main() {
	fmt.Println("Web Scraping...")

	c := colly.NewCollector(
		colly.AllowedDomains("http://yaz.tf.firat.edu.tr", "yaz.tf.firat.edu.tr"),
	)
	var datas []Data
	c.OnHTML(".anno-details", func(e *colly.HTMLElement) {
		selection := e.DOM
		data := Data{
			title:       selection.Find("p.anno-details-title").Text(),
			description: selection.Find("p.anno-details-description").Text(),
		}
		datas = append(datas, data)

	})
	c.Visit("http://yaz.tf.firat.edu.tr/tr/announcements-all")
	fmt.Println(datas[0].description)

}
