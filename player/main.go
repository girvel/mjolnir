package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

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
