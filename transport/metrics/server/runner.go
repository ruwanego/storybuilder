package metrics

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/app/metrics"
)

// Run runs a server to exposes metrics as a separate Prometheus metric server.
func Run(cfg config.AppConfig, _ *container.Container) {
	if !cfg.Metrics.Enabled {
		return
	}

	// register defined metrics
	metrics.Register()

	// set metric exposing port and endpoint
	address := cfg.Host + ":" + strconv.Itoa(cfg.Metrics.Port)
	http.Handle(cfg.Metrics.Route, promhttp.Handler())

	// run metric server in a goroutine so that it doesn't block
	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Println(err)
			panic("Metric server error...")
		}
	}()

	fmt.Printf("Exposing metrics on %v%s ...\n", address, cfg.Metrics.Route)
}
