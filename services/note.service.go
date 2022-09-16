package services

import (
	"net/http"

	"github.com/trixky/tt_orness/endpoints"
)

func StartNoteService() {
	mux := http.NewServeMux()

	mux.HandleFunc("/notes", endpoints.Notes)

	http.ListenAndServe(":3000", mux)
}
