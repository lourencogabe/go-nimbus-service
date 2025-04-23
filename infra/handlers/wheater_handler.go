package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lourencogabe/gonimbus/application/usecases"
)

// responsavel por receber requisiçãos externas para a minha api
func GetWeathersRoute(router *gin.Engine) {
	router.GET("/get/:city", getWeatherHandler)
}

func getWeatherHandler(ctx *gin.Context) {
	city := ctx.Param("city")
	data, err := usecases.GetWeatherData(city)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, data)
}
