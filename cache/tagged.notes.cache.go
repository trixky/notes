package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

type taggedNotesCache struct {
	mtx   sync.Mutex
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
	tn.mtx.Lock()

	notes, ok := tn.Notes[tag]

	if !ok {
		notes = []models.Note{}
	}

	tn.mtx.Unlock()

	return notes
}

func (tn *taggedNotesCache) GetAll() []models.Note {
	tn.mtx.Lock()

	notes := []models.Note{}

	for _, tagged_notes := range tn.Notes {
		notes = append(notes, tagged_notes...)
	}

	tn.mtx.Unlock()

	return notes
}

var TaggedNotes = taggedNotesCache{Notes: make(map[string][]models.Note)}
