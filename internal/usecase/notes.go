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

func (u *noteUsecase) CreateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (*entity.Note, error) {
	noteID, err := u.repository.CreateNote(ctx, payload)
	if err != nil {
		return nil, err
	}

	return u.repository.GetNote(ctx, noteID)
}

func (u *noteUsecase) UpdateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (*entity.Note, error) {
	_, err := u.repository.GetNote(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	noteID, err := u.repository.UpdateNote(ctx, payload)
	if err != nil {
		return nil, err
	}

	return u.repository.GetNote(ctx, noteID)
}

func (u *noteUsecase) DeleteNote(ctx context.Context, noteID string) error {
	_, err := u.repository.GetNote(ctx, noteID)
	if err != nil {
		return err
	}

	return u.repository.DeleteNote(ctx, noteID)
}

func (u *noteUsecase) GetNote(ctx context.Context, noteID string) (*entity.Note, error) {
	return u.repository.GetNote(ctx, noteID)
}

func (u *noteUsecase) GetNoteList(ctx context.Context, params *entity.GetNoteListParams) ([]*entity.Note, int64, error) {
	return u.repository.GetNoteList(ctx, params)
}
