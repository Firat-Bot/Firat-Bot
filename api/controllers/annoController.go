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
	c.OnHTML(".anno-details", func(e *colly.HTMLElement) {
		selection := e.DOM
		event := models.Event{
			Title:       strings.TrimSpace(selection.Find("p.anno-details-title").Text()),
			Description: strings.TrimSpace(selection.Find("p.anno-details-description").Text()),
			Url:         e.ChildAttr("a", "href"),
		}
		events = append(events, event)
	})
	c.Visit("http://yaz.tf.firat.edu.tr/tr/announcements-all")

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, events)
	}
}
