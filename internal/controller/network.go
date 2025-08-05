package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type NetworkController struct {
	networkService service.Network
}

func NewNetworkController() *NetworkController {
	return &NetworkController{}
}

func (n *NetworkController) SetupNetworkService(networkService service.Network) {
	n.networkService = networkService
}

func (n *NetworkController) GetUsageLastFiveSeconds(ctx *gin.Context) {
	data, err := n.networkService.GetUsageLastFiveSeconds()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get network usage"})
		return
	}
	ctx.JSON(200, data)
}
