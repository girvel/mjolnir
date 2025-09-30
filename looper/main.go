package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorhill/cronexpr"
	"github.com/pelletier/go-toml"
)

func logic() error {
    bytes, err := os.ReadFile("sample.toml")
	if err != nil {
	    return err
	}

    var schedule map[string]map[string]string
	err = toml.Unmarshal(bytes, &schedule)
	if err != nil {
	    return err
	}

	for cron_expression, children := range schedule {
	    for id, name := range children {
	        fmt.Println(cronexpr.MustParse(cron_expression).Next(time.Now()), id, name)
	    }
	}
	return nil
}

func index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func main() {
	err := logic()
	if err != nil {
	    panic(err)
	}

    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.Static("/shared", "../shared")

	router.GET("/", index)

	router.Run()
}
