package router

import (
	"net/http"

	"github.com/selahaddinislamoglu/moni/backend/controller"
)

type Router interface {
	SetupCPURoutes(cpuController controller.CPU)
	SetupCORS()
	GetHTTPHandler() (http.Handler, error)
}
