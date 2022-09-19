package validators

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/trixky/tt_orness/config"
	"github.com/trixky/tt_orness/models"
)

func TestNotePost(t *testing.T) {
	tooLong := make([]byte, 1000)
	for i := 0; i < len(tooLong); i++ {
		tooLong[i] = 'a'
	}

	perfectLengthMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH])
	tooLongMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH+1])

	perfectLengthTag := string(tooLong[:config.MAX_TAG_LENGTH])
	tooLongTag := string(tooLong[:config.MAX_TAG_LENGTH+1])

	tests := []struct {
		Body string
		Note models.Note
		Err  error
	}{
		{
			Body: "{\"message\": \"\"}",
			Note: models.Note{
				Message: "",
				Tag:     "",
			},
			Err: models.ErrNoteMessageRequired,
		},
		{
			Body: "{\"message\": \"" + tooLongMessage + "\"}",
			Note: models.Note{
				Message: "",
				Tag:     "",
			},
			Err: models.ErrNoteMessageTooLong,
		},
		{
			Body: "{\"message\": \"" + perfectLengthMessage + "\", \"tag\": \"" + tooLongTag + "\"}",
			Note: models.Note{
				Message: "",
				Tag:     "",
			},
			Err: models.ErrNoteTagTooLong,
		},
		{
			Body: "{\"message\": \"" + perfectLengthMessage + "\", \"tag\": \"" + perfectLengthTag + "\"}",
			Note: models.Note{
				Message: perfectLengthMessage,
				Tag:     perfectLengthTag,
			},
			Err: nil,
		},
		{
			Body: "{\"message\": \"MeSsAgE\"}",
			Note: models.Note{
				Message: "MeSsAgE",
				Tag:     "",
			},
			Err: nil,
		},
	}

	for _, test := range tests {
		bodyReader := strings.NewReader(test.Body)
		bodyReadCloser := io.NopCloser(bodyReader)

		note, err := NotePost(&http.Request{Body: bodyReadCloser})

		if test.Err != nil && (err == nil || !errors.Is(err, test.Err)) {
			t.Fatalf("error expected: %v | result: %v", test.Err, err)
		}
		if test.Err == nil && err != nil {
			t.Fatalf("no error expected: %v | result: %v", test.Err, err)
		}
		if test.Err == nil {
			if test.Note.Message != note.Message {
				t.Fatalf("message don't match: %v | result: %v", test.Note.Message, note.Message)
			}
			if test.Note.Tag != note.Tag {
				t.Fatalf("tag don't match: %v | result: %v", test.Note.Tag, note.Tag)
			}
		}
	}
}
