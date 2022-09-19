package endpoints

import (
	"net/http"

	"github.com/trixky/tt_orness/handlers"
)

// Notes handle the "notes" endpoint.
func Notes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetNotesHandler(w, r)
	case "POST":
		handlers.PostNotesHandler(w, r)
	case "DELETE":
		handlers.DeleteNotesHandler(w, r)
	default:
		// If the method is not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
