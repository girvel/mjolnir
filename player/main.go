package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

const audioPrefix = "static/audio/"

func playlist(c *gin.Context) {
	entries, err := os.ReadDir(audioPrefix)
	if err != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to read audio directory"})
		return
	}

	var items []string
	for _, e := range entries {
		name := e.Name()
		if name[0] == '.' { continue }
	    items = append(items, name)
	}

    c.JSON(http.StatusOK, gin.H{
        "items": items,
		"prefix": audioPrefix,
    })
}

func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", index)
	router.GET("/api/playlist", playlist)

	router.Run()
}
