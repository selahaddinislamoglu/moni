package main

import (
	"github.com/selahaddinislamoglu/moni/internal/controller"
	"github.com/selahaddinislamoglu/moni/internal/router"
	"github.com/selahaddinislamoglu/moni/internal/server"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

func main() {
	secretService := service.NewSecretService()
	authenticationService := service.NewAuthenticationService()
	authenticationService.SetupSecretService(secretService)
	authorizationService := service.NewAuthorizationService()
	authorizationService.SetupSecretService(secretService)
	cpuService := service.NewCPUService()
	memoryService := service.NewMemoryService()
	diskService := service.NewDiskService()
	networkService := service.NewNetworkService()

	authenticationController := controller.NewAuthenticationController()
	authenticationController.SetupAuthenticationService(authenticationService)
	authorizationController := controller.NewAuthorizationController()
	authorizationController.SetupAuthorizationService(authorizationService)
	htmlController := controller.NewHTMLController()
	cpuController := controller.NewCPUController()
	cpuController.SetupCPUService(cpuService)
	memoryController := controller.NewMemoryController()
	memoryController.SetupMemoryService(memoryService)
	diskController := controller.NewDiskController()
	diskController.SetupDiskService(diskService)
	networkController := controller.NewNetworkController()
	networkController.SetupNetworkService(networkService)

	router := router.NewHTTPRouter()
	router.SetupRoutes(authenticationController,
		authorizationController,
		htmlController,
		cpuController,
		memoryController,
		diskController,
		networkController)

	httpServer := server.NewHTTPServer()
	httpServer.SetupRoutes(router)
	if err := httpServer.Serve(":8888"); err != nil {
		panic(err)
	}
}
