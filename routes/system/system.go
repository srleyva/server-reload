package system

import (
	"net/http"

	"github.com/srleyva/server-reload/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// Handler hold the router and logger
type Handler struct {
	*routes.BaseHandler
}

// NewHandler creates a handler for the system
func NewHandler(logger *logrus.Logger) *Handler {
	router := chi.NewRouter()

	handler := &Handler{
		routes.NewBaseHandler(logger, router),
	}

	// System Routes
	router.Get("/health", handler.health)
	router.Handle("/metrics", promhttp.Handler())

	return handler
}

// Health returns the health of the system
func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"ready":   true,
		"version": h.BaseHandler.Conf.Version,
	}

	render.JSON(w, r, health)
}
