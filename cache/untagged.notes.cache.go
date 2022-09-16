package cache

import (
	"sync"
)

type untaggedNotesCache struct {
	mtx   sync.Mutex
	Notes []string
}

func (un *untaggedNotesCache) Add(message string) {
	un.mtx.Lock()

	un.Notes = append(un.Notes, message)

	un.mtx.Unlock()
}

func (un *untaggedNotesCache) Get() []string {
	un.mtx.Lock()

	notes := un.Notes

	un.mtx.Unlock()

	return notes
}

var UntaggedNotes = untaggedNotesCache{Notes: []string{}}
