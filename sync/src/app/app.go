package app

import (
	"net/http"
	"log"
	"time"
	"fmt"
)

type Server struct {
	Configuration ServerConfig
	Handler       ServerHandler
	Responder     ServerResponder
}

type ServerConfig struct {
	LogExecutionTime bool
	Port int
}

type ServerHandler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

type ServerResponder interface {
	InternalServerError(w http.ResponseWriter, r *http.Request)
}

func (s *Server) Load() {
	s.Responder = new(Responder)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if s.Configuration.LogExecutionTime {
		defer LogExecutionTime(time.Now())
	}

	err := s.Handler.Handle(w, r)

	if err != nil {
		log.Printf("Error: %s", err)
		s.Responder.InternalServerError(w, r)
	}
}

func LogExecutionTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Request handled in %s \n", elapsed)
}
