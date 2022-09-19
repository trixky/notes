package validators

import (
	"net/http"

	"github.com/trixky/tt_orness/models"
)

func NoteGet(r *http.Request) (tag string, err error) {
	tag = r.URL.Query().Get("tag")

	note := models.Note{
		Tag: tag,
	}

	if err = note.TagIsValid(); err != nil {
		return
	}

	return
}
