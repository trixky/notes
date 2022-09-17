package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

type untaggedNotesCache struct {
	mtx   sync.RWMutex
	Notes []models.Note
}

func (un *untaggedNotesCache) Add(note models.Note) {
	un.mtx.Lock()

	un.Notes = append(un.Notes, note)

	un.mtx.Unlock()
}

func (un *untaggedNotesCache) Get() []models.Note {
	un.mtx.RLock()

	notes := un.Notes

	un.mtx.RUnlock()

	return notes
}

var UntaggedNotes = untaggedNotesCache{Notes: []models.Note{}}
