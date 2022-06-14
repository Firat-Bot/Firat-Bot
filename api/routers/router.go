package routers

import (
	"github.com/gin-gonic/gin"
	controller "gophers/api/controllers"
)

func AnnoRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/annos", controller.GetAnnos())
}
