package validators

import (
	"net/http"

	"github.com/trixky/tt_orness/models"
)

// NoteGet validates or not the user parameters of the GET request for notes.
func NoteGet(r *http.Request) (tag string, err error) {
	tag = r.URL.Query().Get("tag")

	note := models.Note{
		Tag: tag,
	}

	if err = note.TagIsValid(); err != nil {
		// If the tag is not valid.
		return
	}

	return
}
