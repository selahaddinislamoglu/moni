package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/backend/controller"
)

type HTTPRouter struct {
	cpuController controller.CPU
	ginEngine     *gin.Engine
}

func NewHTTPRouter() Router {
	return &HTTPRouter{
		ginEngine: gin.Default(),
	}
}

func (r *HTTPRouter) SetupCPURoutes(cpuController controller.CPU) {
	r.cpuController = cpuController

	cpu := r.ginEngine.Group("/cpu")
	usage := cpu.Group("/usage")
	usage.GET("/last-five-seconds", r.cpuController.GetUsageLastFiveSeconds)
}

func (r *HTTPRouter) GetHTTPHandler() (http.Handler, error) {
	return r.ginEngine, nil
}

func (r *HTTPRouter) SetupCORS() {
	r.ginEngine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
}

func (r *HTTPRouter) SetupMemoryRoutes(memoryController controller.Memory) {
	memory := r.ginEngine.Group("/memory")
	usage := memory.Group("/usage")
	usage.GET("/all", memoryController.GetUsage)
}

func (r *HTTPRouter) SetupDiskRoutes(diskController controller.Disk) {
	disk := r.ginEngine.Group("/disk")
	usage := disk.Group("/usage")
	usage.GET("/last-five-seconds", diskController.GetUsageLastFiveSeconds)
}
