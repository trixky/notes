package models

import (
	"testing"

	"github.com/trixky/tt_orness/config"
)

func TestIsTaggedNote(t *testing.T) {
	tests := []struct {
		Tag      string
		isTagged bool
	}{
		{
			Tag:      "",
			isTagged: false,
		},
		{
			Tag:      "1",
			isTagged: true,
		},
		{
			Tag:      "12",
			isTagged: true,
		},
		{
			Tag:      "123456789qwer~!@#$%",
			isTagged: true,
		},
	}

	for _, test := range tests {
		note := Note{
			Tag: test.Tag,
		}

		if is_tagged := note.IsTaggedNote(); is_tagged != test.isTagged {
			t.Fatalf("tagged boolean don't match: %v | result: %v", test.Tag, is_tagged)
		}
	}
}

func TestMessageIsValid(t *testing.T) {
	tooLong := make([]byte, 1000)
	for i := 0; i < len(tooLong); i++ {
		tooLong[i] = 'a'
	}

	perfectLengthMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH])
	tooLongMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH+1])

	tests := []struct {
		Message string
		Err     error
	}{
		{
			Message: "",
			Err:     ErrNoteMessageRequired,
		},
		{
			Message: tooLongMessage,
			Err:     ErrNoteMessageTooLong,
		},
		{
			Message: perfectLengthMessage,
			Err:     nil,
		},
		{
			Message: ".",
			Err:     nil,
		},
		{
			Message: "123",
			Err:     nil,
		},
		{
			Message: "123456789qwer~!@#$%",
			Err:     nil,
		},
	}

	for _, test := range tests {
		note := Note{
			Message: test.Message + "",
		}

		if err := note.MessageIsValid(); err != test.Err {
			t.Fatalf("error don't match: %v | result: %v", test.Err, err)
		}
	}
}

func TestTagIsValid(t *testing.T) {
	tooLong := make([]byte, 1000)
	for i := 0; i < len(tooLong); i++ {
		tooLong[i] = 'a'
	}

	perfectLengthTag := string(tooLong[:config.MAX_TAG_LENGTH])
	tooLongTag := string(tooLong[:config.MAX_TAG_LENGTH+1])

	tests := []struct {
		Tag string
		Err error
	}{
		{
			Tag: "",
			Err: nil,
		},
		{
			Tag: tooLongTag,
			Err: ErrNoteTagTooLong,
		},
		{
			Tag: perfectLengthTag,
			Err: nil,
		},
		{
			Tag: ".",
			Err: nil,
		},
		{
			Tag: "123",
			Err: nil,
		},
		{
			Tag: "123456789qwer~!@#$%",
			Err: nil,
		},
	}

	for _, test := range tests {
		note := Note{
			Tag: test.Tag,
		}

		if err := note.TagIsValid(); err != test.Err {
			t.Fatalf("error don't match: %v | result: %v", test.Err, err)
		}
	}
}

func TestIsValid(t *testing.T) {
	tooLong := make([]byte, 1000)
	for i := 0; i < len(tooLong); i++ {
		tooLong[i] = 'a'
	}

	perfectLengthMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH])
	tooLongMessage := string(tooLong[:config.MAX_MESSAGE_LENGTH+1])

	perfectLengthTag := string(tooLong[:config.MAX_TAG_LENGTH])
	tooLongTag := string(tooLong[:config.MAX_TAG_LENGTH+1])

	tests := []struct {
		Message string
		Tag     string
		Err     error
	}{
		{
			Message: "",
			Err:     ErrNoteMessageRequired,
		},
		{
			Message: tooLongMessage,
			Err:     ErrNoteMessageTooLong,
		},
		{
			Message: perfectLengthMessage,
			Err:     nil,
		},
		{
			Message: ".",
			Err:     nil,
		},
		{
			Message: "123",
			Err:     nil,
		},
		{
			Message: "123456789qwer~!@#$%",
			Err:     nil,
		},
		{
			Message: "ok",
			Tag:     "",
			Err:     nil,
		},
		{
			Message: "ok",
			Tag:     tooLongTag,
			Err:     ErrNoteTagTooLong,
		},
		{
			Message: "ok",
			Tag:     perfectLengthTag,
			Err:     nil,
		},
		{
			Message: "ok",
			Tag:     ".",
			Err:     nil,
		},
		{
			Message: "ok",
			Tag:     "123",
			Err:     nil,
		},
		{
			Message: "ok",
			Tag:     "123456789qwer~!@#$%",
			Err:     nil,
		},
		{
			Message: "",
			Tag:     tooLongTag,
			Err:     ErrNoteMessageRequired,
		},
		{
			Message: tooLongMessage,
			Tag:     perfectLengthTag,
			Err:     ErrNoteMessageTooLong,
		},
		{
			Message: perfectLengthMessage,
			Tag:     ".",
			Err:     nil,
		},
	}

	for _, test := range tests {
		note := Note{
			Tag:     test.Tag,
			Message: test.Message,
		}

		if err := note.IsValid(); err != test.Err {
			t.Fatalf("error don't match: %v | result: %v", test.Err, err)
		}
	}
}
