package models

import (
	"errors"
	"strconv"

	"github.com/trixky/tt_orness/config"
)

var ErrNoteMessageRequired = errors.New("message is requried")
var ErrNoteMessageTooLong = errors.New("message is too long (max " + strconv.Itoa(config.MAX_MESSAGE_LENGTH) + " characters)")
var ErrNoteTagTooLong = errors.New("tag is too long (max " + strconv.Itoa(config.MAX_TAG_LENGTH) + " characters)")

type Note struct {
	Message string `json:"message"`
	Tag     string `json:"tag,omitempty"`
}

// IsTaggedNote determines if the note is tagged or not.
func (n *Note) IsTaggedNote() bool {
	return len(n.Tag) > 0
}

// MessageIsValid checks if the message is valid.
func (n *Note) MessageIsValid() error {
	if len(n.Message) == 0 {
		// If the message is absent.
		return ErrNoteMessageRequired
	}

	if len(n.Message) > config.MAX_MESSAGE_LENGTH {
		// If the message length is too long.
		return ErrNoteMessageTooLong
	}

	return nil
}

// MessageIsValid checks if the tag is valid.
func (n *Note) TagIsValid() error {
	if len(n.Tag) > config.MAX_TAG_LENGTH {
		// If the tag length is too long.
		return ErrNoteTagTooLong
	}
	return nil
}

// MessageIsValid checks if all parameters are valid.
func (n *Note) IsValid() error {
	if err := n.MessageIsValid(); err != nil {
		// If the note message is not valid.
		return err
	}

	if err := n.TagIsValid(); err != nil {
		// If the note tag is not valid.
		return err
	}

	return nil
}
