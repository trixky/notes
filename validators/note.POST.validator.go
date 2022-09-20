package validators

import (
	"encoding/json"
	"net/http"

	"github.com/trixky/tt_orness/models"
)

// NotePost validates or not the user parameters of the POST request for notes.
func NotePost(r *http.Request) (new_note models.Note, err error) {
	err = json.NewDecoder(r.Body).Decode(&new_note)

	if err != nil {
		// If an error occurs in the parsing of the body.
		return
	}

	// Check if the not is valid.
	err = new_note.IsValid()

	return
}
