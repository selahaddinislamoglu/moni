package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type CPUController struct {
	cpuService service.CPU
}

func NewCPUController() CPU {
	return &CPUController{}
}

func (c *CPUController) Setup(cpuService service.CPU) {
	c.cpuService = cpuService
}

func (c *CPUController) GetUsageLastFiveSeconds(ctx *gin.Context) {
	data, err := c.cpuService.GetUsageLastFiveSeconds()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get CPU usage"})
		return
	}
	ctx.JSON(200, data)
}
