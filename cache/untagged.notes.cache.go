package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

type untaggedNotesCache struct {
	mtx   sync.Mutex
	Notes []models.Note
}

func (un *untaggedNotesCache) Add(message string) {
	un.mtx.Lock()

	un.Notes = append(un.Notes, models.Note{Message: message})

	un.mtx.Unlock()
}

func (un *untaggedNotesCache) Get() []models.Note {
	un.mtx.Lock()

	notes := un.Notes

	un.mtx.Unlock()

	return notes
}

var UntaggedNotes = untaggedNotesCache{Notes: []models.Note{}}
