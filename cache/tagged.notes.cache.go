package cache

import (
	"sync"
)

type taggedNotesCache struct {
	mtx   sync.Mutex
	Notes map[string][]string
}

func (tn *taggedNotesCache) Add(message, tag string) {
	tn.mtx.Lock()

	if notes, ok := tn.Notes[tag]; ok {
		tn.Notes[tag] = append(notes, message)
	} else {
		tn.Notes[tag] = []string{message}
	}

	tn.mtx.Unlock()
}

func (tn *taggedNotesCache) Get(tag string) []string {
	tn.mtx.Lock()

	notes := tn.Notes[tag]

	tn.mtx.Unlock()

	return notes

}

var TaggedNotes = taggedNotesCache{Notes: make(map[string][]string)}
