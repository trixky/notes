package services

import (
	"net/http"
	"strconv"

	"github.com/trixky/tt_orness/endpoints"
)

func StartNoteService(port int) {
	mux := http.NewServeMux()

	mux.HandleFunc("/notes", endpoints.Notes)

	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}
