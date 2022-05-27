package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/transport/http/router"
)

// Run runs the http server.
func Run(cfg config.AppConfig, ctr *container.Container) *http.Server {
	// initialize the router
	r := router.Init(ctr)
	srv := &http.Server{
		Addr: cfg.Host + ":" + strconv.Itoa(cfg.Port),
		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: cfg.Timeout.Write.Dur(),
		ReadTimeout:  cfg.Timeout.Read.Dur(),
		IdleTimeout:  cfg.Timeout.Idle.Dur(),
		// pass our instance of chi.Mux in
		Handler: r,
	}
	// run our server in a goroutine so that it doesn't block
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
			panic("Service shutting down unexpectedly...")
		}
	}()
	fmt.Println("Service started...")
	fmt.Printf("Listening on %v ...\n", srv.Addr)
	return srv
}

// Stop stops the server.
func Stop(ctx context.Context, srv *http.Server) {
	fmt.Println("Service shutting down...")
	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Error Shutting the server down: %v", err)
	}
}
