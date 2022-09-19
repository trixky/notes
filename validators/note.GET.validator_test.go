package validators

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/trixky/tt_orness/config"
	"github.com/trixky/tt_orness/models"
)

func TestNoteGet(t *testing.T) {
	tooLong := make([]byte, 1000)
	for i := 0; i < len(tooLong); i++ {
		tooLong[i] = 'a'
	}

	perfectLengthTag := string(tooLong[:config.MAX_TAG_LENGTH])
	tooLongTag := string(tooLong[:config.MAX_TAG_LENGTH+1])

	tests := []struct {
		URL url.URL
		Tag string
		Err error
	}{
		{
			URL: url.URL{
				RawQuery: "tag=" + tooLongTag,
			},
			Tag: "",
			Err: models.ErrNoteTagTooLong,
		},
		{
			URL: url.URL{},
			Tag: "",
			Err: nil,
		},
		{
			URL: url.URL{
				RawQuery: "tag=",
			},
			Tag: "",
			Err: nil,
		},
		{
			URL: url.URL{
				RawQuery: "tagg=" + perfectLengthTag,
			},
			Tag: "",
			Err: nil,
		},
		{
			URL: url.URL{
				RawQuery: "tag=" + perfectLengthTag,
			},
			Tag: perfectLengthTag,
			Err: nil,
		},
		{
			URL: url.URL{
				RawQuery: "tag=TaG",
			},
			Tag: "TaG",
			Err: nil,
		},
	}

	for _, test := range tests {

		tag, err := NoteGet(&http.Request{URL: &test.URL})

		if test.Err != nil && (err == nil || !errors.Is(err, test.Err)) {
			t.Fatalf("error expected: %v | result: %v", test.Err, err)
		}
		if test.Err == nil && err != nil {
			t.Fatalf("no error expected: %v | result: %v", test.Err, err)
		}
		if test.Err == nil {
			if test.Tag != tag {
				t.Fatalf("tag don't match: %v | result: %v", test.Tag, tag)
			}
		}
	}
}
