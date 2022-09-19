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

func (n *Note) IsTaggedNote() bool {
	return len(n.Tag) > 0
}

func (n *Note) MessageIsValid() error {
	if len(n.Message) == 0 {
		return ErrNoteMessageRequired
	}

	if len(n.Message) > config.MAX_MESSAGE_LENGTH {
		return ErrNoteMessageTooLong
	}

	return nil
}

func (n *Note) TagIsValid() error {
	if len(n.Tag) > config.MAX_TAG_LENGTH {
		return ErrNoteTagTooLong
	}
	return nil
}

func (n *Note) IsValid() error {
	if err := n.MessageIsValid(); err != nil {
		return err
	}

	if err := n.TagIsValid(); err != nil {
		return err
	}

	return nil
}
