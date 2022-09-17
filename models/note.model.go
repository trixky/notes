package models

type Note struct {
	Message string `json:"message"`
	Tag     string `json:"tag"`
}

func (n *Note) IsTaggedNote() bool {
	return len(n.Tag) > 0
}
