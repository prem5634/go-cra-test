// Command tasksvc runs a small, in-memory task-tracking HTTP API.
package main

import (
	"log"
	"net/http"
	"time"

	"gitlab.com/prembito/go-cra-test/api"
	"gitlab.com/prembito/go-cra-test/auth"
	"gitlab.com/prembito/go-cra-test/config"
	"gitlab.com/prembito/go-cra-test/store"
)

func main() {
	cfg, err := config.Load("config/app.yaml")
	if err != nil {
		log.Printf("config: falling back to defaults: %v", err)
	}

	h := &api.Handler{
		Store:  store.New(),
		Issuer: auth.NewIssuer(time.Hour),
	}

	srv := &http.Server{
		Addr:        cfg.Addr,
		Handler:     api.NewRouter(h, cfg.LogRequests),
		ReadTimeout: time.Duration(cfg.ReadTimeout) * time.Second,
	}

	log.Printf("tasksvc listening on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
