package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

type taggedNotesCache struct {
	mtx   sync.RWMutex
	Notes map[string][]models.Note
}

func (tn *taggedNotesCache) Add(note models.Note) {
	tn.mtx.Lock()

	if notes, ok := tn.Notes[note.Tag]; ok {
		tn.Notes[note.Tag] = append(notes, note)
	} else {
		tn.Notes[note.Tag] = []models.Note{note}
	}

	tn.mtx.Unlock()
}

func (tn *taggedNotesCache) Get(tag string) []models.Note {
	tn.mtx.RLock()

	notes, ok := tn.Notes[tag]

	if !ok {
		notes = []models.Note{}
	}

	tn.mtx.RUnlock()

	return notes
}

func (tn *taggedNotesCache) GetAll() []models.Note {
	tn.mtx.RLock()

	notes := []models.Note{}

	for _, tagged_notes := range tn.Notes {
		notes = append(notes, tagged_notes...)
	}

	tn.mtx.RUnlock()

	return notes
}

func (tn *taggedNotesCache) Delete(tag string) (number_of_deleted_notes int) {
	tn.mtx.Lock()

	if notes, ok := tn.Notes[tag]; ok {
		number_of_deleted_notes = len(notes)
		delete(tn.Notes, tag)
	}

	tn.mtx.Unlock()

	return
}

func (tn *taggedNotesCache) DeleteAll() (number_of_deleted_notes int) {
	tn.mtx.Lock()

	for tag := range tn.Notes {
		number_of_deleted_notes += len(tn.Notes[tag])
		delete(tn.Notes, tag)
	}

	tn.mtx.Unlock()

	return
}

var TaggedNotes = taggedNotesCache{Notes: make(map[string][]models.Note)}
