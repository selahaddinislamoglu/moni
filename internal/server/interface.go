package server

import "github.com/selahaddinislamoglu/moni/internal/router"

type Server interface {
	Setup(router router.Router)
	Serve(addr string) error
	Shutdown() error
}
