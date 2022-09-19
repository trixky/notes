package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

// untaggedNotesCache saves untagged notes in the cache (not persistent).
type untaggedNotesCache struct {
	mtx   sync.RWMutex
	Notes []models.Note
}

// Add adds untagged a note in the cache.
func (un *untaggedNotesCache) Add(note models.Note) {
	un.mtx.Lock()

	// Add the note.
	un.Notes = append(un.Notes, note)

	un.mtx.Unlock()
}

// Get gets all untagged notes in the cache.
func (un *untaggedNotesCache) Get() []models.Note {
	un.mtx.RLock()

	// Get all notes.
	notes := un.Notes

	un.mtx.RUnlock()

	return notes
}

// DeleteAll deletes all untagged notes in the cache.
func (un *untaggedNotesCache) DeleteAll() (number_of_deleted_notes int) {
	un.mtx.Lock()

	// Get the number of notes that will be deleted.
	number_of_deleted_notes = len(un.Notes)
	// Delete all notes.
	un.Notes = nil

	un.mtx.Unlock()

	return
}

// Init the global "untagged notes" cache.
var UntaggedNotes = untaggedNotesCache{Notes: []models.Note{}}
