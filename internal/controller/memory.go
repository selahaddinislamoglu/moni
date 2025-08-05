package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type MemoryController struct {
	memoryService service.Memory
}

func NewMemoryController() Memory {
	return &MemoryController{}
}

func (c *MemoryController) SetupMemoryService(memoryService service.Memory) {
	c.memoryService = memoryService
}

func (c *MemoryController) GetUsage(ctx *gin.Context) {
	data, err := c.memoryService.GetUsage()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get memory usage"})
		return
	}
	ctx.JSON(200, data)
}
