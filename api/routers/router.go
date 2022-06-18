package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gophers/api/controllers"
	"net/http"
	"os"
)

func AnnoRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/annos", controller.GetAnnos())
}

func Router() gin.IRoutes {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	AnnoRoutes(router)

	fmt.Println(port, ": connected")
	router.Run(port)

	url := "http://localhost:8080/annos"
	return Response(url, router)

}
func Response(url string, router *gin.Engine) gin.IRoutes {

	resp, _ := http.Get(url)
	if resp.StatusCode == 200 {
		return router.Use()
	}
	return router.GET("/annos", controller.GetAnnos())
}