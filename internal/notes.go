package internal

import (
	"context"
	"notes-api/internal/entity"
)

type NoteUsecase interface {
	CreateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (*entity.Note, error)
	UpdateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (*entity.Note, error)
	DeleteNote(ctx context.Context, noteID string) error
	GetNote(ctx context.Context, noteID string) (*entity.Note, error)
	GetNoteList(ctx context.Context, params *entity.GetNoteListParams) ([]*entity.Note, int64, error)
}

type NoteRepository interface {
	CreateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (string, error)
	UpdateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (string, error)
	DeleteNote(ctx context.Context, noteID string) error
	GetNote(ctx context.Context, noteID string) (*entity.Note, error)
	GetNoteList(ctx context.Context, params *entity.GetNoteListParams) ([]*entity.Note, int64, error)
}
