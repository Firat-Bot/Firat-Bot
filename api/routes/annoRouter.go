package router

import (
	controller "gophers/api/controllers"

	"github.com/gin-gonic/gin"
)

func AnnoRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/annos", controller.GetAnnos())
}
