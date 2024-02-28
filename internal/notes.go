package internal

import (
	"context"
	"notes-api/internal/entity"
)

type NoteUsecase interface {
	CreateNote(ctx context.Context, payload *entity.Note) (*entity.Note, error)
	UpdateNote(ctx context.Context, payload *entity.Note) (*entity.Note, error)
	DeleteNote(ctx context.Context, bookID string) error
	GetNote(ctx context.Context, bookID string) (*entity.Note, error)
	GetNoteList(ctx context.Context, filter *entity.GetNoteListFilter) ([]*entity.Note, int64, error)
}

type NoteRepository interface {
	CreateNote(ctx context.Context, payload *entity.Note) (string, error)
	UpdateNote(ctx context.Context, payload *entity.Note) (string, error)
	DeleteNote(ctx context.Context, bookID string) error
	GetNote(ctx context.Context, bookID string) (*entity.Note, error)
	GetNoteList(ctx context.Context, filter *entity.GetNoteListFilter) ([]*entity.Note, int64, error)
}
