package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

func get_tasks() ([]string, error) {
    bytes, err := os.ReadFile("sample.toml")
	if err != nil {
	    return nil, err
	}

    var schedule map[string]map[string]string
	err = toml.Unmarshal(bytes, &schedule)
	if err != nil {
	    return nil, err
	}

	var result []string
	for _, children := range schedule {
	    for _, name := range children {
			result = append(result, name)
	    }
	}
	return result, nil
}

func index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func tasks_get(c *gin.Context) {
    tasks, err := get_tasks()
	if err != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.Static("/shared", "../shared")

	router.GET("/", index)
	router.GET("/tasks/", tasks_get)

	router.Run()
}
