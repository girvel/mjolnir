package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/girvel/mjolnir/api/docs"
	api "github.com/girvel/mjolnir/api/src"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API for local IoT network
// @description

// @host thor1:8080

func main() {
    router := gin.Default()
	router.GET("/sensors", api.Sensors)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
