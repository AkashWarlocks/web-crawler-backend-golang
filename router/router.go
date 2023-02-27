package routes

import (
	controller "example/scrapper/controller"

	"github.com/gin-gonic/gin"
)

func RouteIndex(Router *gin.Engine) {
	Router.GET("/crypto", controller.CryptoData)
}