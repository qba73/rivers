package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	river "github.com/qba73/rivers"
	"github.com/qba73/rivers/cmd/rivers-api/handlers"
)

var (
	addr string
)

func ListStations(w http.ResponseWriter, r *http.Request) error {
	stations, err := river.LoadStations("latesttest.json")
	if err != nil {
		return err
	}
	riverStations := stations.All()

	output, err := json.Marshal(&riverStations)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return nil
}

func main() {
	flag.StringVar(&addr, "addr", ":5000", ":5000")
	flag.Parse()

	logger := log.New(os.Stdout, "RIVERS-API", log.LstdFlags)

	ih := handlers.NewInfo(logger)
	sh := handlers.NewStations(logger)

	// ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", ih)
	mux.Handle("/stations", sh)

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

	// wait for either interrupt or kill signals and shutdown server gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// block and wait for the signal
	sig := <-c
	logger.Println("received signal:", sig)

	// shutdown server gracefully but wait for 20 sec for current requests to complete
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	server.Shutdown(ctx)
}
