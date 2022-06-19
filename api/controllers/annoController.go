package controller

import (
	"net/http"
	"strings"

	"gophers/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func GetAnnos() gin.HandlerFunc {

	var events []models.Event
	c := colly.NewCollector(
		colly.AllowedDomains("http://yaz.tf.firat.edu.tr", "yaz.tf.firat.edu.tr"),
	)
	c.OnHTML(".anno-details", func(element *colly.HTMLElement) {
		selection := element.DOM
		event := models.Event{
			Title:       strings.TrimSpace(selection.Find("p.anno-details-title").Text()),
			Description: strings.TrimSpace(selection.Find("p.anno-details-description").Text()),
			Url:         element.ChildAttr("a", "href"),
		}
		events = append(events, event)
	})
	c.Visit("http://yaz.tf.firat.edu.tr/tr/announcements-all")

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, events)
	}
}
func GetInfoForLecturer() gin.HandlerFunc {

	var lecturer []models.Lecturer

	c := colly.NewCollector(
		colly.AllowedDomains("http://yaz.tf.firat.edu.tr", "yaz.tf.firat.edu.tr"),
	)
	//var name string
	c.OnHTML(".staff-content-text", func(element *colly.HTMLElement) {

		var all []string
		all = append(all, element.ChildText("p"))
		splitEPosta := strings.Split(all[0], "Telefon")
		splitPhone := strings.Split(splitEPosta[1], "Çalışma")
		if strings.Contains(splitEPosta[1], "Kişisel") {
			splitPhone = strings.Split(splitEPosta[1], "Kişisel")
		}

		lecture := models.Lecturer{
			Name:      element.ChildText("h3"),
			Email:     splitEPosta[0],
			Phone:     splitPhone[0],
			Workspace: splitPhone[1],
		}
		lecturer = append(lecturer, lecture)

	})
	c.Visit("http://yaz.tf.firat.edu.tr/tr/academic-staffs")

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, lecturer)
	}
}
