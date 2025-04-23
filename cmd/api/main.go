package main

import (
	"github.com/gin-gonic/gin"
	infra "github.com/lourencogabe/gonimbus/infra/handlers"
)

func main() {
	router := gin.Default()
	infra.GetWeathersRoute(router)
	router.Run("localhost:8080")
}
