package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @title  thor1 homepage
// @description  Website with information about the home server

// @host  thor1

// @Summary  Homepage
// @Description  Provides an overview of the server, contains all general information about its capabilities and local network
// @Produce  html
// @Router  / [get]
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H {})
}

func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", index)

	router.Run()
}
