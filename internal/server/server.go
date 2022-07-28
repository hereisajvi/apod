package server

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/chiefcake/apod/internal/config"
)

// Server contains methods for serving and shutting down the server.
type Server struct {
	server *http.Server
}

// New configures routes and returns a server instance.
func New(cfg *config.Config, handler PictureHandler) *Server {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").
		Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").
		Subrouter()

	v1Router.HandleFunc("/pictures", handler.List).
		Methods(http.MethodGet)

	v1Router.HandleFunc("/pictures/{date}", handler.GetByDate).
		Methods(http.MethodGet)

	return &Server{
		server: &http.Server{
			Addr:    net.JoinHostPort(cfg.ServerHost, cfg.ServerPort),
			Handler: router,
		},
	}
}

// Serve starts initialized server.
func (s Server) Serve() error {
	listener, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return errors.Wrap(err, "could not build tcp listener")
	}

	err = s.server.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "could not serve gateway server")
	}

	return nil
}

// Shutdown stops initialized server with provided context.
func (s Server) Shutdown(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "could not shutdown gateway server")
	}

	return nil
}
