package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Data struct {
	title       string
	description string
	url         string
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
			title:       strings.TrimSpace(selection.Find("p.anno-details-title").Text()),
			description: strings.TrimSpace(selection.Find("p.anno-details-description").Text()),
			url:         e.ChildAttr("a", "href"),
		}
		fmt.Println(data)

		datas = append(datas, data)

	})
	c.Visit("http://yaz.tf.firat.edu.tr/tr/announcements-all")
	file, err := os.Create("announcements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, i := range datas {
		_, err := file.WriteString(i.title)
		_, err = file.WriteString(i.description)
		_, err = file.WriteString(i.url)
		if err != nil {
			log.Fatal(err)
		}
	}
}
