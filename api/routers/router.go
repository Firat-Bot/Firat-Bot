package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gophers/api/controllers"
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
	return router.GET("/annos", controller.GetAnnos())

}
