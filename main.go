package main

import (
	"fmt"
	"os"

	routers "gophers/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Web Scraping...")

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	routers.AnnoRoutes(router)

	fmt.Println(port, ": connected")
	router.Run(port)
}
