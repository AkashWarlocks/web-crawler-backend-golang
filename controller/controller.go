package controller

import (
	utils "example/scrapper/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CryptoData(c *gin.Context) {
	resposneData  := utils.GetCryptodata()
	c.JSON(http.StatusOK, resposneData)

}

func GetFile(c *gin.Context) {

}