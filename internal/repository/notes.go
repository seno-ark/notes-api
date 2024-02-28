package repository

import "notes-api/internal"

type noteRepository struct{}

func NewNoteRepository() internal.NoteRepository {
	return &noteRepository{}
}
