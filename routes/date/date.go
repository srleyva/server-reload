package date

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/srleyva/server-reload/routes"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

// Handler hold the router and logger
type Handler struct {
	*routes.BaseHandler
}

func NewHandler(logger *logrus.Logger) *Handler {

	router := chi.NewRouter()

	handler := &Handler{
		routes.NewBaseHandler(logger, router),
	}

	// Date Routes
	router.Get("/", handler.date)
	return handler
}

func (h *Handler) date(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r,
		map[string]interface{}{
			"time": time.Now(),
		},
	)
}
