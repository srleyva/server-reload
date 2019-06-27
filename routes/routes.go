package routes

import (
	"github.com/go-chi/chi"
	"github.com/srleyva/server-reload/pkg/config"
)

type Handler interface {
	Router() *chi.Mux
	SetConfig(*config.ServerConfiguration)
}
