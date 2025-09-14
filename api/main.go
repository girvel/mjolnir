package main

import (
	"context"
	"os"
	"log/slog"

	"github.com/gin-gonic/gin"
	_ "github.com/girvel/mjolnir/api/docs"
	api "github.com/girvel/mjolnir/api/src"
	"github.com/jackc/pgx/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API for local IoT network
// @description

// @host thor1:8080

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@postgres:5432/home")
	if err != nil {
	    slog.Error("Unable to connect to database", "err", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	controller := api.Controller{Conn: conn};

    router := gin.Default()
	router.GET("/sensors", controller.Sensors)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
