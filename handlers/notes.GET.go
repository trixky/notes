package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/models"
	"github.com/trixky/tt_orness/validators"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	tag, err := validators.NoteGet(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	var notes []models.Note

	if len(tag) > 0 {
		notes = cache.TaggedNotes.Get(tag)
	} else {
		notes = cache.TaggedNotes.GetAll()
		notes = append(notes, cache.UntaggedNotes.Get()...)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(notes)
}
