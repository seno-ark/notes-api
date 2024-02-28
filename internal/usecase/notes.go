package usecase

import (
	"context"
	"notes-api/internal"
	"notes-api/internal/entity"
	appErr "notes-api/pkg/error"
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
	noteID, err := u.repository.CreateNote(ctx, payload)
	if err != nil {
		return nil, appErr.NewErrInternalServer(err.Error())
	}

	note, err := u.repository.GetNote(ctx, noteID)
	if err != nil {
		return nil, appErr.NewErrInternalServer(err.Error())
	}

	return note, nil
}

func (u *noteUsecase) UpdateNote(ctx context.Context, payload *entity.Note) (*entity.Note, error) {
	_, err := u.repository.GetNote(ctx, payload.ID)
	if err != nil {
		if appErr.IsErrNotFound(err) {
			return nil, appErr.NewErrNotFound("note not found")
		}
		return nil, appErr.NewErrInternalServer(err.Error())
	}

	noteID, err := u.repository.UpdateNote(ctx, payload)
	if err != nil {
		return nil, appErr.NewErrInternalServer(err.Error())
	}

	note, err := u.repository.GetNote(ctx, noteID)
	if err != nil {
		return nil, appErr.NewErrInternalServer(err.Error())
	}

	return note, nil
}

func (u *noteUsecase) DeleteNote(ctx context.Context, noteID string) error {
	_, err := u.repository.GetNote(ctx, noteID)
	if err != nil {
		if appErr.IsErrNotFound(err) {
			return appErr.NewErrNotFound("note not found")
		}
		return appErr.NewErrInternalServer(err.Error())
	}

	err = u.repository.DeleteNote(ctx, noteID)
	if err != nil {
		return appErr.NewErrInternalServer(err.Error())
	}

	return nil
}

func (u *noteUsecase) GetNote(ctx context.Context, noteID string) (*entity.Note, error) {
	note, err := u.repository.GetNote(ctx, noteID)
	if err != nil {
		if appErr.IsErrNotFound(err) {
			return nil, appErr.NewErrNotFound("note not found")
		}
		return nil, appErr.NewErrInternalServer(err.Error())
	}

	return note, nil
}

func (u *noteUsecase) GetNoteList(ctx context.Context, filter *entity.GetNoteListFilter) ([]*entity.Note, int64, error) {
	notes, total, err := u.repository.GetNoteList(ctx, filter)
	if err != nil {
		return nil, total, appErr.NewErrInternalServer(err.Error())
	}

	return notes, total, nil
}
