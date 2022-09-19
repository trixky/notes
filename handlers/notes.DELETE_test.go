package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/models"
)

func TestDeleteNotesHandler(t *testing.T) {
	const target = "http://localhost:3000/notes"
	const tag = "TaG"

	tests := []struct {
		Notes          []models.Note
		ExpectedStatus int
		Tagged         int
	}{
		{
			Notes: []models.Note{
				{
					Message: "cat",
				},
			},
			ExpectedStatus: 200,
			Tagged:         0,
		},
		{
			Notes: []models.Note{
				{
					Message: "cat",
					Tag:     tag,
				},
			},
			ExpectedStatus: 200,
			Tagged:         1,
		},
		{
			Notes: []models.Note{
				{
					Message: "cat",
				},
				{
					Message: "cat",
					Tag:     tag,
				},
			},
			ExpectedStatus: 200,
			Tagged:         1,
		},
		{
			Notes: []models.Note{
				{
					Message: "cat",
				},
				{
					Message: "cat",
					Tag:     tag,
				},
				{
					Message: "cat",
				},
				{
					Message: "cat",
					Tag:     tag,
				},
			},
			ExpectedStatus: 200,
			Tagged:         2,
		},
	}

	for _, _tag := range []string{"", "?tag=" + tag} {
		for _, test := range tests {
			cache.TaggedNotes.DeleteAll()
			cache.UntaggedNotes.DeleteAll()

			for _, note := range test.Notes {
				if note.IsTaggedNote() {
					cache.TaggedNotes.Add(note)
				} else {
					cache.UntaggedNotes.Add(note)
				}
			}

			req := httptest.NewRequest("GET", target+_tag, nil)
			w := httptest.NewRecorder()
			DeleteNotesHandler(w, req)

			resp := w.Result()

			if resp.StatusCode != test.ExpectedStatus {
				t.Fatalf("status code don't match: %v | result: %v", test.ExpectedStatus, resp.StatusCode)
			}

			if test.ExpectedStatus == http.StatusOK {
				var len_test_notes int
				if len(_tag) > 0 {
					len_test_notes = len(test.Notes) - test.Tagged
				} else {
					len_test_notes = 0
				}

				len_cached_notes := len(cache.TaggedNotes.Notes) + len(cache.UntaggedNotes.Notes)

				if len_test_notes != len_cached_notes {
					t.Fatalf("length of cached messages don't match: %v | result: %v", len_test_notes, len_cached_notes)
				}
			}
		}
	}
}
