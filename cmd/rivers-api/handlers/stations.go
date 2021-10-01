package handlers

import (
	"log"
	"net/http"

	"github.com/qba73/rivers/cmd/rivers-api/data"
)

type Stations struct {
	l *log.Logger
}

func NewStations(l *log.Logger) Stations {
	return Stations{l}
}

func (s Stations) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sx := data.GetStations()
	if err := sx.ToJSON(w); err != nil {
		http.Error(w, "error mashaling json", http.StatusInternalServerError)
	}
}
