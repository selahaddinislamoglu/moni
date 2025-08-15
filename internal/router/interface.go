package router

import (
	"net/http"

	"github.com/selahaddinislamoglu/moni/internal/controller"
)

type Router interface {
	SetupRoutes(authenticationController controller.Authentication,
		authorizationController controller.Authorization,
		websocketController controller.Websocket,
		htmlController controller.HTML,
		cpuController controller.CPU,
		memoryController controller.Memory,
		diskController controller.Disk,
		networkController controller.Network)
	GetHTTPHandler() (http.Handler, error)
}
