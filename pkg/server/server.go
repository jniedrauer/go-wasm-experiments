package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

// Start starts the webserver
func Start() {
	handler := handlers.LoggingHandler(os.Stdout, getRouter())

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
