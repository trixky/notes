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

func (un *untaggedNotesCache) DeleteAll() (number_of_deleted_notes int) {
	un.mtx.Lock()

	number_of_deleted_notes = len(un.Notes)
	un.Notes = nil

	un.mtx.Unlock()

	return
}

var UntaggedNotes = untaggedNotesCache{Notes: []models.Note{}}
