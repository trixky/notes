package services

import (
	"net/http"
	"strconv"

	"github.com/trixky/tt_orness/endpoints"
)

// StartNoteService starts the "note" service
func StartNoteService(port int) {
	mux := http.NewServeMux()

	mux.HandleFunc("/notes", endpoints.Notes)

	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}
