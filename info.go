package rivers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ====================
// Data

// InfoVersion it's metadata about the rivers API.
type InfoVersion struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	ReleasedOn string `json:"released_on,omitempty"`
}

// ToJSON knows how to encode info version into json.
func (iv InfoVersion) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(iv)
}

// NewInfoVersion knows how to create a new version.
func NewInfoVersion() *InfoVersion {
	iv := InfoVersion{
		Name:       "rivers",
		Version:    "v0.1.0",
		ReleasedOn: "",
	}
	return &iv
}

// ====================
// Handlers

type InfoVersionHandler struct {
	l *log.Logger
}

func NewVersionHandler(l *log.Logger) *InfoVersionHandler {
	return &InfoVersionHandler{l}
}

func (i *InfoVersionHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	iv := NewInfoVersion()
	if err := iv.ToJSON(w); err != nil {
		http.Error(w, "Unable to marshall json", http.StatusInternalServerError)
	}
}
