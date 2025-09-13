package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorsOutput struct {
    Temperature float32 `json:"temperature"`
    Humidity float32 `json:"humidity"`
}

// @Summary Request temperature & humidity values from sensors
// @Description
// @Produce json
// @Success 200 {object} api.SensorsOutput
// @Router /sensors [get]
func Sensors(c *gin.Context) {
	c.JSON(http.StatusOK, SensorsOutput{
		Temperature: 25.0,
		Humidity: 52.0,
	})
}

