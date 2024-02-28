package usecase

import (
	"context"
	"notes-api/internal"
	"notes-api/internal/entity"
)

type noteUsecase struct {
	repository internal.NoteRepository
}

func NewNoteUsecase(repository internal.NoteRepository) internal.NoteUsecase {
	return &noteUsecase{
		repository: repository,
	}
}

func (u *noteUsecase) CreateNote(ctx context.Context, payload *entity.Note) (*entity.Note, error) {
	return nil, nil
}

func (u *noteUsecase) UpdateNote(ctx context.Context, payload *entity.Note) (*entity.Note, error) {
	return nil, nil
}

func (u *noteUsecase) DeleteNote(ctx context.Context, bookID string) error {
	return nil
}

func (u *noteUsecase) GetNote(ctx context.Context, bookID string) (*entity.Note, error) {
	return nil, nil
}

func (u *noteUsecase) GetNoteList(ctx context.Context, filter *entity.GetNoteListFilter) ([]*entity.Note, int64, error) {
	return nil, 0, nil
}
