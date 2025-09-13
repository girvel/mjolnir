package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/girvel/mjolnir/homepage/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title thor1 homepage
// @description Website with information about the home server

// @host thor1

// @Summary Homepage
// @Description Provides an overview of the server
// @Produce html
// @Router / [get]
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H {})
}

func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", index)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
