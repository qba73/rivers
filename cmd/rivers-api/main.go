package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	river "github.com/qba73/rivers"
)

var (
	addr     string
	certfile string
	keyfile  string
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"name": "Rivers project API", "version": "v1"}`)
	//w.Write([]byte(`{"name": "Rivers project API", "version": "v1"}`))
}

func handlerListStations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := ListStations(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
	flag.StringVar(&certfile, "cert", "", "server certificate file")
	flag.StringVar(&keyfile, "key", "", "server key file")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/stations", handlerListStations).Methods(http.MethodGet)

	log.Println("starting server on :5000")

	server := http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServeTLS(certfile, keyfile); err != nil {
		log.Fatal(err)
	}
}
