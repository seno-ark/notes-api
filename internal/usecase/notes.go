package usecase

import "notes-api/internal"

type noteUsecase struct {
	repository internal.NoteRepository
}

func NewNoteUsecase(repository internal.NoteRepository) internal.NoteUsecase {
	return &noteUsecase{
		repository: repository,
	}
}
