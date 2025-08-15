package main

import (
	"encoding/json"
	"fmt"

	"github.com/selahaddinislamoglu/moni/internal/controller"
	"github.com/selahaddinislamoglu/moni/internal/router"
	"github.com/selahaddinislamoglu/moni/internal/server"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

func main() {

	brokerService := service.NewBrokerService()
	loggerSubscriber := service.NewSubscriberService()
	loggerSubscriber.Setup(brokerService)
	loggerFunc := func(topic string) func(json.RawMessage) {
		return func(message json.RawMessage) {
			fmt.Printf("[%s] %s\n", topic, string(message))
		}
	}
	loggerSubscriber.Subscribe(service.CPUTopic, loggerFunc(service.CPUTopic))
	loggerSubscriber.Subscribe(service.MemoryTopic, loggerFunc(service.MemoryTopic))
	loggerSubscriber.Subscribe(service.DiskTopic, loggerFunc(service.DiskTopic))
	loggerSubscriber.Subscribe(service.NetworkTopic, loggerFunc(service.NetworkTopic))

	secretService := service.NewSecretService()
	authenticationService := service.NewAuthenticationService()
	authenticationService.Setup(secretService)
	authorizationService := service.NewAuthorizationService()
	authorizationService.Setup(secretService)
	websocketService := service.NewWebsocketService()
	websocketService.Setup(brokerService)

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
	authenticationController.Setup(authenticationService)
	authorizationController := controller.NewAuthorizationController()
	authorizationController.Setup(authorizationService)
	websocketController := controller.NewWebsocketController()
	websocketController.Setup(websocketService)
	htmlController := controller.NewHTMLController()

	cpuController := controller.NewCPUController()
	cpuController.Setup(cpuService)
	memoryController := controller.NewMemoryController()
	memoryController.Setup(memoryService)
	diskController := controller.NewDiskController()
	diskController.Setup(diskService)
	networkController := controller.NewNetworkController()
	networkController.Setup(networkService)

	router := router.NewHTTPRouter()
	router.SetupRoutes(authenticationController,
		authorizationController,
		websocketController,
		htmlController,
		cpuController,
		memoryController,
		diskController,
		networkController)

	httpServer := server.NewHTTPServer()
	httpServer.Setup(router)
	if err := httpServer.Serve(":8888"); err != nil {
		panic(err)
	}
}
