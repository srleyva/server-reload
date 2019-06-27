package routes

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/srleyva/server-reload/pkg/config"
)

type Handler interface {
	Router() *chi.Mux
	SetConfig(*config.ServerConfiguration)
}

type BaseHandler struct {
	Conf   *config.ServerConfiguration
	router *chi.Mux
	logger *logrus.Logger
}

func NewBaseHandler(logger *logrus.Logger, router *chi.Mux) *BaseHandler {
	return &BaseHandler{
		logger: logger,
		router: router,
	}
}

// SetConfig set the config for this module
func (b *BaseHandler) SetConfig(conf *config.ServerConfiguration) {
	b.Conf = conf
}

// Router returns the router
func (b *BaseHandler) Router() *chi.Mux {
	return b.router
}
