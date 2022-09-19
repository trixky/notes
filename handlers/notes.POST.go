package handlers

import (
	"fmt"
	"net/http"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/validators"
)

// PostNotesHandler handle the POST methode for the "notes" endpoint.
func PostNotesHandler(w http.ResponseWriter, r *http.Request) {
	new_note, err := validators.NotePost(r)

	if err != nil {
		// If the request inputs are not valid.
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	if new_note.IsTaggedNote() {
		// If a tag was specified.
		cache.TaggedNotes.Add(new_note)
	} else {
		// If no tag was specified.
		cache.UntaggedNotes.Add(new_note)
	}
}
