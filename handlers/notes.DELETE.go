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

func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	tag, err := validators.NoteDelete(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	var number_of_deleted_notes = 0

	if len(tag) > 0 {
		number_of_deleted_notes = cache.TaggedNotes.Delete(tag)
	} else {
		number_of_deleted_notes = cache.TaggedNotes.DeleteAll()
		number_of_deleted_notes += cache.UntaggedNotes.DeleteAll()
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(DeleteResponse{
		NumberOfDeletedNotes: number_of_deleted_notes,
	})
}
