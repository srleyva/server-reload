package main

import (
	"os"
	"os/signal"

	"github.com/srleyva/server-reload/routes"

	"github.com/srleyva/server-reload/pkg/config"
	"github.com/srleyva/server-reload/pkg/server"

	"github.com/srleyva/server-reload/routes/date"
	"github.com/srleyva/server-reload/routes/system"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

// VERSION is passed in as a flag
var (
	VERSION = "0.0.0"
)

func main() {
	logger := logrus.New()
	v1, err := config.ReadConfig(".env", nil)
	if err != nil {
		logger.Fatalf("Error reading in config: %s", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	reload := make(chan string)
	v1.OnConfigChange(func(e fsnotify.Event) {
		reload <- "reload"
	})

	v1.Set("server.version", VERSION)

	routes := map[string]routes.Handler{
		"/date":   date.NewHandler(logger),
		"/system": system.NewHandler(logger),
	}

	server.Start(quit, reload, logger, v1, "/api", routes)

}
