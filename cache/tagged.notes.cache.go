package cache

import (
	"sync"

	"github.com/trixky/tt_orness/models"
)

type taggedNotesCache struct {
	mtx   sync.Mutex
	Notes map[string][]models.Note
}

func (tn *taggedNotesCache) Add(message, tag string) {
	tn.mtx.Lock()

	if notes, ok := tn.Notes[tag]; ok {
		tn.Notes[tag] = append(notes, models.Note{Message: message})
	} else {
		tn.Notes[tag] = []models.Note{{Message: message}}
	}

	tn.mtx.Unlock()
}

func (tn *taggedNotesCache) Get(tag string) []models.Note {
	tn.mtx.Lock()

	notes := tn.Notes[tag]

	tn.mtx.Unlock()

	return notes

}

var TaggedNotes = taggedNotesCache{Notes: make(map[string][]models.Note)}
