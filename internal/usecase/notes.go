package usecase

import "notes-api/internal"

type noteUsecase struct {
	repository internal.NoteUsecase
}

func NewNoteUsecase(repository internal.NoteUsecase) internal.NoteUsecase {
	return &noteUsecase{
		repository: repository,
	}
}
