package main

import (
	"fmt"

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

	brokerService := service.NewBrokerService()
	loggerSubscriber := service.NewSubscriberService()
	loggerSubscriber.Setup(brokerService)
	loggerFunc := func(topic string) func([]byte) {
		return func(message []byte) {
			fmt.Printf("[%s] %s\n", topic, string(message))
		}
	}
	loggerSubscriber.Subscribe(service.CPUTopic, loggerFunc(service.CPUTopic))
	loggerSubscriber.Subscribe(service.MemoryTopic, loggerFunc(service.MemoryTopic))
	loggerSubscriber.Subscribe(service.DiskTopic, loggerFunc(service.DiskTopic))
	loggerSubscriber.Subscribe(service.NetworkTopic, loggerFunc(service.NetworkTopic))

	cpuPublisher := service.NewPublisherService()
	cpuPublisher.Setup(brokerService)
	cpuService := service.NewCPUService()
	cpuService.Setup(cpuPublisher)

	memoryPublisher := service.NewPublisherService()
	memoryPublisher.Setup(brokerService)
	memoryService := service.NewMemoryService()
	memoryService.Setup(memoryPublisher)

	diskPublisher := service.NewPublisherService()
	diskPublisher.Setup(brokerService)
	diskService := service.NewDiskService()
	diskService.Setup(diskPublisher)

	networkPublisher := service.NewPublisherService()
	networkPublisher.Setup(brokerService)
	networkService := service.NewNetworkService()
	networkService.Setup(networkPublisher)

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
