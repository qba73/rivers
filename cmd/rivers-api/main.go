package main

import (
	"flag"
	"log"
	"net/http"
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
	w.Write([]byte(`{"name": "Rivers project API", "version": "v1"}`))
}

func showFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"type": "Feature", "properties": {"name": "Sandy Mills", "ref": "0000001041"}, "geometry": {"type": "Point", "coordinates": [-7.575758, 54.838318]}}`))
}

func showStation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific gauge station"))
}


func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/feature", showFeature)
	mux.HandleFunc("/station", showStation)

	log.Println("starting server on :5000")

	server := http.Server{
		Addr: addr,
		Handler: mux,
	}

	if err := server.ListenAndServeTLS(certfile, keyfile); err != nil {
		log.Fatal(err)
	}
}
