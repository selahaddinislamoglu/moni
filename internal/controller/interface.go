package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type CPU interface {
	SetupCPUService(cpuService service.CPU)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Memory interface {
	SetupMemoryService(memoryService service.Memory)
	GetUsage(ctx *gin.Context)
}

type Disk interface {
	SetupDiskService(diskService service.Disk)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Network interface {
	SetupNetworkService(networkService service.Network)
	GetUsageLastFiveSeconds(ctx *gin.Context)
}

type Authentication interface {
	SetupAuthenticationService(authenticationService service.Authentication)
	Login(ctx *gin.Context)
}

type Authorization interface {
	SetupAuthorizationService(authorizationService service.Authorization)
	IsAuthorized(ctx *gin.Context) bool
}

type HTML interface {
	Login(ctx *gin.Context)
	Dashboard(ctx *gin.Context)
}
