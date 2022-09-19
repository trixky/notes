package cache

import (
	"math"
	"strconv"
	"testing"

	"github.com/trixky/tt_orness/models"
)

func TestAddTagged(t *testing.T) {
	taggedNotes := taggedNotesCache{Notes: make(map[string][]models.Note)}

	for i := 1; i < 1000; i++ {
		taggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
			Tag:     "TaG" + strconv.Itoa(i%100),
		})

		total_len := 0

		for tag := range taggedNotes.Notes {
			total_len += len(taggedNotes.Notes[tag])
		}

		if total_len != i {
			t.Fatalf("number of cached notes don't match: %v | result: %v", total_len, i)
		}
	}
}

func TestGetTagged(t *testing.T) {
	taggedNotes := taggedNotesCache{Notes: make(map[string][]models.Note)}

	for i := 0; i < 1000; i++ {
		tag := "TaG" + strconv.Itoa(i%100)
		taggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
			Tag:     tag,
		})

		total_len := len(taggedNotes.Get(tag))
		expected_len := int(int(math.Floor(float64(i/100)))) + 1

		if total_len != expected_len {
			t.Fatalf("number of getted notes don't match: %v | result: %v", total_len, expected_len)
		}
	}
}

func TestGetAllTagged(t *testing.T) {
	taggedNotes := taggedNotesCache{Notes: make(map[string][]models.Note)}

	for i := 0; i < 1000; i++ {
		tag := "TaG" + strconv.Itoa(i%100)
		taggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
			Tag:     tag,
		})

		total_len := len(taggedNotes.GetAll())
		expected_len := i + 1

		if total_len != expected_len {
			t.Fatalf("number of getted notes don't match: %v | result: %v", total_len, expected_len)
		}
	}
}

func TestDeleteTagged(t *testing.T) {
	taggedNotes := taggedNotesCache{Notes: make(map[string][]models.Note)}

	for i := 0; i < 1000; i++ {
		tag := "TaG" + strconv.Itoa(i%100)
		taggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
			Tag:     tag,
		})
	}

	for i := 100; i > 0; i-- {
		tag := "TaG" + strconv.Itoa(i)
		taggedNotes.Delete(tag)

		total_len := len(taggedNotes.Get(tag))
		if total_len != 0 {
			t.Fatalf("number of deleted notes don't match: %v | result: %v", total_len, 0)
		}

		total_all_len := len(taggedNotes.GetAll())
		expected_len := i * 10

		if total_all_len != expected_len {
			t.Fatalf("number of all deleted notes don't match: %v | result: %v", total_all_len, expected_len)
		}
	}
}

func TestDeleteAllTagged(t *testing.T) {
	taggedNotes := taggedNotesCache{Notes: make(map[string][]models.Note)}

	for i := 0; i < 1000; i++ {
		tag := "TaG" + strconv.Itoa(i%100)
		taggedNotes.Add(models.Note{
			Message: "MeSsAgE" + strconv.Itoa(i%100),
			Tag:     tag,
		})
	}

	taggedNotes.DeleteAll()

	total_len := len(taggedNotes.Notes)

	if total_len != 0 {
		t.Fatalf("number of deleted notes don't match: %v | result: %v", total_len, 0)
	}
}
