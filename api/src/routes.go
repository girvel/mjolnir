package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
    Conn *pgx.Conn
}

type SensorsOutput struct {
    Temperature float32 `json:"temperature"`
    Humidity float32 `json:"humidity"`
}

// @Summary Request temperature & humidity values from sensors
// @Description
// @Produce json
// @Success 200 {object} api.SensorsOutput
// @Router /sensors [get]
func (co *Controller) Sensors(c *gin.Context) {
	var humidity float32
	co.Conn.QueryRow(context.Background(), `
		SELECT value
		FROM humidity
		ORDER BY timestamp DESC
	`).Scan(&humidity)

	var temperature float32
	co.Conn.QueryRow(context.Background(), `
		SELECT value
		FROM temperature
		ORDER BY timestamp DESC
	`).Scan(&temperature)

	c.JSON(http.StatusOK, SensorsOutput{
		Temperature: temperature,
		Humidity: humidity,
	})
}

