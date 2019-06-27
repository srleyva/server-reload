package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	chilogger "github.com/766b/chi-logger"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/srleyva/server-reload/pkg/config"
	"github.com/srleyva/server-reload/routes"
)

// Start Handles the running of the server and configuring it
// This includes functionality to reload on the fly
func Start(quit chan os.Signal, reload chan string, logger *logrus.Logger, v1 *viper.Viper, rootRoute string, routes map[string]routes.Handler) {

	// For reload unmarshall
	var conf config.Configuration
	if err := v1.Unmarshal(&conf); err != nil {
		logrus.Fatalf("err reading conf: %s", err)
	}

	logger.SetFormatter(&logrus.JSONFormatter{})
	if conf.Logging.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	r := chi.NewRouter()
	r.Use(chilogger.NewLogrusMiddleware("router", logger))

	root := r.Route(rootRoute, func(r chi.Router) {
		for k, v := range routes {
			v.SetConfig(&conf.Server)
			r.Mount(k, v.Router())
		}
	})

	// Enable root list of routes
	if conf.Server.List {
		root.Get("/", func(w http.ResponseWriter, r *http.Request) {
			logger.Warning("Listing enabled on root")
			routes := make(map[string][]string)
			walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
				routes[route] = append(routes[route], method)
				return nil
			}
			if err := chi.Walk(root, walkFunc); err != nil {
				logger.Fatalf("err walking routes: %s", err)
			}

			json.NewEncoder(w).Encode(&routes)
		})
	}

	port := conf.Server.Port
	hostname := conf.Server.Hostname
	version := conf.Server.Version
	addr := fmt.Sprintf("%s:%d", hostname, port)

	h := http.Server{
		Addr:    addr,
		Handler: r,
	}
	go h.ListenAndServe()
	logger.Infof("Server Version %s running on %s", version, addr)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		select {
		case <-quit:
			logger.Fatal("Shutting Down")
			h.Shutdown(ctx)
		case <-reload:
			logger.Info("Reloading the config")
			h.Shutdown(ctx)
			go Start(quit, reload, logger, v1, rootRoute, routes)
		}
	}
}
