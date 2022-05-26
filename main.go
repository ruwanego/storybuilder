package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/app/splash"
	httpServer "github.com/storybuilder/storybuilder/transport/http/server"
	metricsServer "github.com/storybuilder/storybuilder/transport/metrics/server"
)

func main() {
	// show splash screen when starting
	splash.Show(splash.StyleDefault)

	// parse all configurations
	cfg := config.Parse("./configs")
	// prints all parsed configurations
	cfg.Show()

	// resolve the container using parsed configurations
	ctr := container.Resolve(cfg)

	// start the server to handle http requests
	srv := httpServer.Run(cfg.AppConfig, ctr)

	// start the server to expose application metrics
	metricsServer.Run(cfg.AppConfig, ctr)

	// enable graceful shutdown
	c := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)

	// block until a registered signal is received
	<-c

	// create a deadline to wait for
	var wait time.Duration

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// release resources
	ctr.Destruct()

	// gracefully stop the http server
	httpServer.Stop(ctx, srv)

	os.Exit(0)
}
