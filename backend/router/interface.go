package router

import (
	"net/http"

	"github.com/selahaddinislamoglu/moni/backend/controller"
)

type Router interface {
	SetupCPURoutes(cpuController controller.CPU)
	SetupMemoryRoutes(memoryController controller.Memory)
	SetupCORS()
	GetHTTPHandler() (http.Handler, error)
}
