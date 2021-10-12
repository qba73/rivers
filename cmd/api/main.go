package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/qba73/rivers"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", ":5000", ":5000")
	flag.Parse()

	logger := log.New(os.Stdout, "RIVERS-API ", log.LstdFlags)

	ih := rivers.NewVersionHandler(logger)
	sh := rivers.NewStationsHandler(logger)

	// ServeMux
	mux := mux.NewRouter()

	// Inforamtion about API version
	mux.HandleFunc("/info", ih.GetVersion).Methods(http.MethodGet)
	mux.HandleFunc("/stations", sh.GetStations).Methods(http.MethodGet)

	// Server
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ErrorLog:     logger,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Printf("staring server on port %s\n", addr)

		if err := server.ListenAndServe(); err != nil {
			logger.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for signals to shutdown the server gracefully.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block and wait for the signal.
	sig := <-c
	logger.Println("received signal:", sig)

	// shutdown server gracefully but wait for 20 sec for current requests to complete
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	server.Shutdown(ctx)
}
