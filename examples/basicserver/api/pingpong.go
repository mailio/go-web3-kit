package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mailio/go-web3-kit/examples/basicserver/models"
)

// Pong method
// @Summary Returns pong message
// @Description Returns all available virtual machien types and their resource capacities
// @Tags PONG API
// @Success 200 {object} models.Pingpong
// @Accept json
// @Produce json
// @Router /api/pong [get]
func PingPongAPI(c *gin.Context) {
	// Public API Definitions
	c.JSON(http.StatusOK, models.Pingpong{Message: "pong"})
}
