package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/controller"
)

type HTTPRouter struct {
	authenticationController controller.Authentication
	authorizationController  controller.Authorization
	htmlController           controller.HTML
	cpuController            controller.CPU
	memoryController         controller.Memory
	diskController           controller.Disk
	networkController        controller.Network
	ginEngine                *gin.Engine
}

func NewHTTPRouter() Router {
	return &HTTPRouter{
		ginEngine: gin.Default(),
	}
}

func (r *HTTPRouter) GetHTTPHandler() (http.Handler, error) {
	return r.ginEngine, nil
}

func (r *HTTPRouter) SetupRoutes(authenticationController controller.Authentication,
	authorizationController controller.Authorization,
	websocketController controller.Websocket,
	htmlController controller.HTML,
	cpuController controller.CPU,
	memoryController controller.Memory,
	diskController controller.Disk,
	networkController controller.Network) {
	r.authenticationController = authenticationController
	r.authorizationController = authorizationController
	r.htmlController = htmlController
	r.cpuController = cpuController
	r.memoryController = memoryController
	r.diskController = diskController
	r.networkController = networkController

	public := r.ginEngine.Group("/")
	public.GET("/login", r.htmlController.Login)
	public.GET("/", r.htmlController.Login)
	public.POST("/login", r.authenticationController.Login)
	public.GET("/dashboard", r.htmlController.Dashboard)

	api := r.ginEngine.Group("/api")
	api.Use(func(c *gin.Context) {
		if !r.authorizationController.IsAuthorized(c) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	})
	api.GET("/cpu/usage/last-five-seconds", r.cpuController.GetUsageLastFiveSeconds)
	api.GET("/memory/usage/all", r.memoryController.GetUsage)
	api.GET("/disk/usage/last-five-seconds", r.diskController.GetUsageLastFiveSeconds)
	api.GET("/network/usage/last-five-seconds", r.networkController.GetUsageLastFiveSeconds)
	api.GET("/connect", websocketController.Connect)
}
