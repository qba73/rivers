package rivers

import (
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	lg *log.Logger
}

func NewInfo(lg *log.Logger) *Info {
	return &Info{lg}
}

func (i *Info) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.lg.Println("handle info request")
	fmt.Fprintf(w, "Rivers version: %s", "0.1.0")
}
