package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"utfpr_db/dto"
)

type healthCheckController struct {
}

func NewHealthCheckController() *healthCheckController {
	return &healthCheckController{}
}

// HealthCheck - health-check for the server
// @Summary - Health-Check
// @Description - Health-Check for the API
// @Tags - Health-Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthCheckResponse
// @Router /health-check [get]
func (u *healthCheckController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, &dto.HealthCheckResponse{
		Status: "OK",
	})
}
