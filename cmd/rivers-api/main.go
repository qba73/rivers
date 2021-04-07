package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/qba73/rivers/internal/river"
)

var (
	addr string
	certfile string
	keyfile string
)

func init()  {
	flag.StringVar(&addr, "addr", ":5000", ":5000")
	flag.StringVar(&certfile, "cert", "", "server certificate file")
	flag.StringVar(&keyfile, "key", "", "server key file")
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"name": "Rivers project API", "version": "v1"}`)
	//w.Write([]byte(`{"name": "Rivers project API", "version": "v1"}`))
}

func handleStations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed" , http.StatusMethodNotAllowed)
		return
	}
	if err := stations(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func stations(w http.ResponseWriter, r *http.Request) error {
	stations, err := river.LoadStations("latesttest.json")
	if err != nil {
		return err
	}
	riverStations := stations.GetAll()
	output, err := json.Marshal(&riverStations)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return nil
}


func showFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"type": "Feature", "properties": {"name": "Sandy Mills", "ref": "0000001041"}, "geometry": {"type": "Point", "coordinates": [-7.575758, 54.838318]}}`))
}



func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/feature", showFeature).Methods(http.MethodGet)
	r.HandleFunc("/stations", handleStations).Methods(http.MethodGet)



	log.Println("starting server on :5000")

	server := http.Server{
		Addr: addr,
		Handler: r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	if err := server.ListenAndServeTLS(certfile, keyfile); err != nil {
		log.Fatal(err)
	}
}


