package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/trixky/tt_orness/cache"
	"github.com/trixky/tt_orness/config"
	"github.com/trixky/tt_orness/models"
)

func TestPostNotesHandler(t *testing.T) {
	const target = "http://localhost:3000/notes"

	tooLong := make([]byte, 1000)
	for i := 0; i < len(tooLong); i++ {
		tooLong[i] = 'a'
	}

	perfectLengthMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH])
	tooLongMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH+1])

	perfectLengthTag := string(tooLong[:config.MAX_TAG_LENGTH])
	tooLongTag := string(tooLong[:config.MAX_TAG_LENGTH+1])

	tests := []struct {
		Body           string
		Note           models.Note
		ExpectedStatus int
		Tagged         bool
	}{
		{
			Body:           "{}",
			ExpectedStatus: 400,
			Tagged:         false,
		},
		{
			Body:           "{\"message\": \"\"}",
			ExpectedStatus: 400,
			Tagged:         false,
		},
		{
			Body: "{\"message\": \"MeSsAgE\"}",
			Note: models.Note{
				Message: "MeSsAgE",
				Tag:     "",
			},
			ExpectedStatus: 200,
			Tagged:         false,
		},
		{
			Body: "{\"message\": \"" + perfectLengthMessage + "\"}",
			Note: models.Note{
				Message: perfectLengthMessage,
				Tag:     "",
			},
			ExpectedStatus: 200,
			Tagged:         false,
		},
		{
			Body:           "{\"message\": \"" + tooLongMessage + "\"}",
			ExpectedStatus: 400,
			Tagged:         false,
		},
		{
			Body: "{\"message\": \"MeSsAgE\", \"tag\": \"" + perfectLengthTag + "\"}",

			Note: models.Note{
				Message: "MeSsAgE",
				Tag:     perfectLengthTag,
			},
			ExpectedStatus: 200,
			Tagged:         true,
		},
		{
			Body:           "{\"message\": \"MeSsAgE\", \"tag\": \"" + tooLongTag + "\"}",
			ExpectedStatus: 400,
			Tagged:         true,
		},
		{
			Body: "{\"message\": \"MeSsAgE\", \"tag\": \"TaG\"}",
			Note: models.Note{
				Message: "MeSsAgE",
				Tag:     "TaG",
			},
			ExpectedStatus: 200,
			Tagged:         true,
		},
	}

	total_added := 0

	cache.TaggedNotes.DeleteAll()
	cache.UntaggedNotes.DeleteAll()

	for _, test := range tests {
		bodyReader := strings.NewReader(test.Body)

		req := httptest.NewRequest("POST", target, bodyReader)
		w := httptest.NewRecorder()
		PostNotesHandler(w, req)

		resp := w.Result()

		if resp.StatusCode != test.ExpectedStatus {
			t.Fatalf("status code don't match: %v | result: %v", test.ExpectedStatus, resp.StatusCode)
		}

		if test.ExpectedStatus == http.StatusOK {
			total_added++

			var notes []models.Note
			if test.Tagged {
				notes = cache.TaggedNotes.Get(test.Note.Tag)
			} else {
				notes = cache.UntaggedNotes.Get()
			}

			total_cache_len := len(cache.TaggedNotes.Notes) + len(cache.UntaggedNotes.Notes)

			if total_added != total_cache_len {
				t.Fatalf("number of added notes don't match: %v | result: %v", total_added, total_cache_len)
			}

			last_note := notes[len(notes)-1]

			if test.Note.Message != last_note.Message {
				t.Fatalf("message don't match: %v | result: %v", test.Note.Message, last_note.Message)
			}
			if test.Note.Tag != last_note.Tag {
				t.Fatalf("tag don't match: %v | result: %v", test.Note.Tag, last_note.Tag)
			}
		}
	}
}
