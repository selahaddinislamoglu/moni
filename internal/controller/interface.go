package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type CPU interface {
	Setup(cpuService service.CPU)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Memory interface {
	Setup(memoryService service.Memory)
	GetUsage(ctx *gin.Context)
}

type Disk interface {
	Setup(diskService service.Disk)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Network interface {
	Setup(networkService service.Network)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Authentication interface {
	Setup(authenticationService service.Authentication)
	Login(ctx *gin.Context)
}

type Authorization interface {
	Setup(authorizationService service.Authorization)
	IsAuthorized(ctx *gin.Context) bool
}

type HTML interface {
	Login(ctx *gin.Context)
	Dashboard(ctx *gin.Context)
}
