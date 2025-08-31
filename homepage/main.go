package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func (c *gin.Context) {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H {})
	})

	router.Run()
}
