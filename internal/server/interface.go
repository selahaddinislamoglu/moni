package server

import "github.com/selahaddinislamoglu/moni/internal/router"

type Server interface {
	SetupRoutes(router router.Router)
	Serve(addr string) error
	Shutdown() error
}
