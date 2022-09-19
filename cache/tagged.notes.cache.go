package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

// taggedNotesCache saves tagged notes in the cache (not persistent).
type taggedNotesCache struct {
	mtx   sync.RWMutex
	Notes map[string][]models.Note
}

// Add adds a tagged note in the cache.
func (tn *taggedNotesCache) Add(note models.Note) {
	tn.mtx.Lock()

	if notes, ok := tn.Notes[note.Tag]; ok {
		// If the tag is already known
		// Add the note with others associated notes.
		tn.Notes[note.Tag] = append(notes, note)
	} else {
		// If the tag is new
		// Create a new group for notes.
		tn.Notes[note.Tag] = []models.Note{note}
	}

	tn.mtx.Unlock()
}

// Get gets tagged notes in the cache for a specific tag.
func (tn *taggedNotesCache) Get(tag string) []models.Note {
	tn.mtx.RLock()

	// Get note for this tag.
	notes, ok := tn.Notes[tag]

	if !ok {
		// If no known note for this tag.
		notes = []models.Note{}
	}

	tn.mtx.RUnlock()

	return notes
}

// GetAll gets all tagged notes in the cache.
func (tn *taggedNotesCache) GetAll() []models.Note {
	tn.mtx.RLock()

	notes := []models.Note{}

	for _, tagged_notes := range tn.Notes {
		// For each tag known in the cache
		// Get all notes associated with this tag.
		notes = append(notes, tagged_notes...)
	}

	tn.mtx.RUnlock()

	return notes
}

// Delete deletes tagged notes in the cache for a specific tag.
func (tn *taggedNotesCache) Delete(tag string) (number_of_deleted_notes int) {
	tn.mtx.Lock()

	if notes, ok := tn.Notes[tag]; ok {
		// If the tag is known.
		// Get the number of notes that will be deleted.
		number_of_deleted_notes = len(notes)
		// Delete all notes associated with this tag.
		delete(tn.Notes, tag)
	}

	tn.mtx.Unlock()

	return
}

// DeleteAll deletes all tagged notes in the cache.
func (tn *taggedNotesCache) DeleteAll() (number_of_deleted_notes int) {
	tn.mtx.Lock()

	for tag := range tn.Notes {
		// For each tag known in the cache.
		// Get the number of notes that will be deleted.
		number_of_deleted_notes += len(tn.Notes[tag])
		// Delete all notes associated with this tag.
		delete(tn.Notes, tag)
	}

	tn.mtx.Unlock()

	return
}

// Init the global "tagged notes" cache.
var TaggedNotes = taggedNotesCache{Notes: make(map[string][]models.Note)}
