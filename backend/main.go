package main

import (
	"github.com/selahaddinislamoglu/moni/backend/controller"
	"github.com/selahaddinislamoglu/moni/backend/router"
	"github.com/selahaddinislamoglu/moni/backend/server"
	"github.com/selahaddinislamoglu/moni/backend/service"
)

func main() {
	cpuService := service.NewCPUService()
	memoryService := service.NewMemoryService()
	diskService := service.NewDiskService()
	networkService := service.NewNetworkService()

	cpuController := controller.NewCPUController()
	cpuController.SetupCPUService(cpuService)

	memoryController := controller.NewMemoryController()
	memoryController.SetupMemoryService(memoryService)

	diskController := controller.NewDiskController()
	diskController.SetupDiskService(diskService)

	networkController := controller.NewNetworkController()
	networkController.SetupNetworkService(networkService)

	router := router.NewHTTPRouter()
	router.SetupCORS()
	router.SetupCPURoutes(cpuController)
	router.SetupMemoryRoutes(memoryController)
	router.SetupDiskRoutes(diskController)
	router.SetupNetworkRoutes(networkController)

	httpServer := server.NewHTTPServer()
	httpServer.SetupRoutes(router)
	if err := httpServer.Serve(":8080"); err != nil {
		panic(err)
	}
}
