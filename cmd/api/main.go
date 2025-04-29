package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	infra "github.com/lourencogabe/gonimbus/infra/handlers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	infra.GetWeathersRoute(router)
	router.Run("localhost:8080")
}
