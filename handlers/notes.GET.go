package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/models"
	"github.com/trixky/tt_orness/validators"
)

// GetNotesHandler handle the GET methode for the "notes" endpoint.
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	tag, err := validators.NoteGet(r)

	if err != nil {
		// If the request inputs are not valid.
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	var notes []models.Note

	if len(tag) > 0 {
		// If a tag was specified.
		notes = cache.TaggedNotes.Get(tag)
	} else {
		// If no tag was specified.
		notes = cache.TaggedNotes.GetAll()
		notes = append(notes, cache.UntaggedNotes.Get()...)
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the response.
	json.NewEncoder(w).Encode(notes)
}
