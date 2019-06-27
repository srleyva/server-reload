package date

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/srleyva/server-reload/pkg/config"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

// Handler hold the router and logger
type Handler struct {
	conf   *config.ServerConfiguration // TODO: This is nil
	router *chi.Mux
	logger *logrus.Logger
}

func NewHandler(logger *logrus.Logger) *Handler {

	router := chi.NewRouter()
	handler := &Handler{
		logger: logger,
		router: router,
	}

	// Date Routes
	router.Get("/", handler.date)
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

func (h *Handler) date(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r,
		map[string]interface{}{
			"time": time.Now(),
		},
	)
}
