package server

import (
	"net/http"

	"github.com/selahaddinislamoglu/moni/backend/router"
)

type HTTPServer struct {
	router router.Router
}

func NewHTTPServer() Server {
	return &HTTPServer{}
}

func (s *HTTPServer) SetupRoutes(router router.Router) {
	s.router = router
}

func (s *HTTPServer) Serve(addr string) error {
	handler, err := s.router.GetHTTPHandler()
	if err != nil {
		return err
	}

	return http.ListenAndServe(addr, handler)
}

func (s *HTTPServer) Shutdown() error {
	return nil
}
