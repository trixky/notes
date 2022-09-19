package handlers

import (
	"fmt"
	"net/http"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/validators"
)

func PostNotesHandler(w http.ResponseWriter, r *http.Request) {
	new_note, err := validators.NotePost(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	if new_note.IsTaggedNote() {
		cache.TaggedNotes.Add(new_note)
	} else {
		cache.UntaggedNotes.Add(new_note)
	}
}
