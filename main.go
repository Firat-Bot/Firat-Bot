package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Data struct {
	titles      string
	description string
}
type Datas struct {
	datas []Datas
}

func main() {
	fmt.Println("Web Scraping...")

	c := colly.NewCollector(
		colly.AllowedDomains("yaz.tf.firat.edu.tr"),
	)

	c.OnHTML(".announcements", func(e *colly.HTMLElement) {

		links := e.ChildAttrs("p", "class")
		fmt.Println(links, "\n")

	})
	c.Visit("http://yaz.tf.firat.edu.tr/tr/announcements-all")

}
