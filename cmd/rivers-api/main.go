package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not alowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name": "Rivers project API", "version": "v1"}`))
}

func showFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not alowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"type": "Feature", "properties": {"name": "Sandy Mills", "ref": "0000001041"}, "geometry": {"type": "Point", "coordinates": [-7.575758, 54.838318]}}`))
}

func showStation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific gauge station"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/feature", showFeature)
	mux.HandleFunc("/station", showStation)

	log.Println("starting server on :5000")
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatal(err)
	}
}
