package transport

import (
	"fmt"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"net/http"
)

type HttpServer struct {
	router interfaces.IRouter
	Server *http.Server
}

type Listener struct {
	router interfaces.IRouter
}

func (l Listener) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	connection := Connection{
		request:  r,
		response: w,
	}
	l.router.IndicateRoute(connection)
	return
}

func (s *HttpServer) Start(port interfaces.HttpPortNumber) error {
	router := s.router
	if router == nil {
		return fmt.Errorf("router is not defined - cannot start")
	}
	s.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: Listener{router},
	}
	s.Server.ListenAndServe()
	return nil
}