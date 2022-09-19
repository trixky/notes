package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/validators"
)

type DeleteResponse struct {
	NumberOfDeletedNotes int `json:"number_of_deleted_notes"`
}

// DeleteNotesHandler handle the DELETE methode for the "notes" endpoint.
func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	tag, err := validators.NoteDelete(r)

	if err != nil {
		// If the request inputs are not valid.
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	var number_of_deleted_notes = 0

	if len(tag) > 0 {
		// If a tag was specified.
		number_of_deleted_notes = cache.TaggedNotes.Delete(tag)
	} else {
		// If no tag was specified.
		number_of_deleted_notes = cache.TaggedNotes.DeleteAll()
		number_of_deleted_notes += cache.UntaggedNotes.DeleteAll()
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the response.
	json.NewEncoder(w).Encode(DeleteResponse{
		NumberOfDeletedNotes: number_of_deleted_notes,
	})
}
