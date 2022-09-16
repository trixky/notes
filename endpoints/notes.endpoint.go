package endpoints

import (
	"net/http"

	"github.com/trixky/tt_orness/handlers"
)

func Notes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetNotesHandler(w, r)
	case "POST":
		handlers.PostNotesHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
