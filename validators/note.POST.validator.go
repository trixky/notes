package validators

import (
	"encoding/json"
	"net/http"

	"github.com/trixky/tt_orness/models"
)

func NotePost(r *http.Request) (new_note models.Note, err error) {
	err = json.NewDecoder(r.Body).Decode(&new_note)

	if err != nil {
		return
	}

	err = new_note.IsValid()

	return
}
