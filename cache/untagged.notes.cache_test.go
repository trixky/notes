package cache

import (
	"strconv"
	"testing"

	"github.com/trixky/tt_orness/models"
)

func TestAddUntagged(t *testing.T) {
	untaggedNotes := untaggedNotesCache{Notes: []models.Note{}}

	for i := 1; i < 1000; i++ {
		untaggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
		})

		total_len := 0

		total_len += len(untaggedNotes.Notes)

		if total_len != i {
			t.Fatalf("number of cached notes don't match: %v | result: %v", total_len, i)
		}
	}
}

func TestGetUntagged(t *testing.T) {
	untaggedNotes := untaggedNotesCache{Notes: []models.Note{}}

	for i := 0; i < 1000; i++ {
		untaggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
		})

		total_len := len(untaggedNotes.Get())
		expected_len := i + 1

		if total_len != expected_len {
			t.Fatalf("number of getted notes don't match: %v | result: %v", total_len, expected_len)
		}
	}
}

func TestDeleteAllUntagged(t *testing.T) {
	untaggedNotes := untaggedNotesCache{Notes: []models.Note{}}

	for i := 0; i < 1000; i++ {
		untaggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
		})
	}

	untaggedNotes.DeleteAll()

	total_len := len(untaggedNotes.Get())
	expected_len := 0

	if total_len != expected_len {
		t.Fatalf("number of staying notes don't match: %v | result: %v", total_len, expected_len)
	}
}
