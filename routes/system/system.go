package system

import (
	"net/http"

	"github.com/srleyva/server-reload/pkg/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// Handler hold the router and logger
type Handler struct {
	conf   *config.ServerConfiguration
	router *chi.Mux
	logger *logrus.Logger
}

// NewHandler creates a handler for the system
func NewHandler(logger *logrus.Logger) *Handler {

	router := chi.NewRouter()

	handler := &Handler{
		logger: logger,
		router: router,
	}

	// System Routes
	router.Get("/health", handler.health)
	router.Handle("/metrics", promhttp.Handler())

	return handler
}

// SetConfig set the config for this module
func (h *Handler) SetConfig(conf *config.ServerConfiguration) {
	h.conf = conf
}

// Router returns the router
func (h *Handler) Router() *chi.Mux {
	return h.router
}

// Health returns the health of the system
func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"ready":   true,
		"version": h.conf.Version,
	}

	render.JSON(w, r, health)
}
