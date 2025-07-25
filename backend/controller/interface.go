package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/backend/service"
)

type CPU interface {
	SetupCPUService(cpuService service.CPU)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Memory interface {
	SetupMemoryService(memoryService service.Memory)
	GetUsage(ctx *gin.Context)
}
