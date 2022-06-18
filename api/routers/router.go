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
func LecturesRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/lectures", controller.GetInfoForLecturer())

}

func RouterAnnounce() gin.IRoutes {
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
	return Response(url, router, "annos")

}

func RouterLectures() gin.IRoutes {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	LecturesRoute(router)
	fmt.Println(port, ": connected")
	router.Run(port)

	url := "http://localhost:8080/lectures"
	return Response(url, router, "lectures")

}
func Response(url string, router *gin.Engine, path string) gin.IRoutes {

	resp, _ := http.Get(url)
	if resp.StatusCode == 200 {
		return router.Use()
	}
	if path == "annos" {
		return router.GET(path, controller.GetAnnos())

	} else if path == "lectures" {
		return router.GET(path, controller.GetInfoForLecturer())
	}
	return nil
}
