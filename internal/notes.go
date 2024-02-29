// Package internal provides the application core bussines
package internal

import (
	"context"
	"notes-api/internal/entity"
)

// NoteUsecase is usecase abstaction for note
type NoteUsecase interface {
	CreateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (*entity.Note, error)
	UpdateNote(ctx context.Context, noteID string, payload *entity.CreateUpdateNotePayload) (*entity.Note, error)
	DeleteNote(ctx context.Context, noteID string) error
	GetNote(ctx context.Context, noteID string) (*entity.Note, error)
	GetNoteList(ctx context.Context, params *entity.GetNoteListParams) ([]*entity.Note, int64, error)
}

// NoteRepository is repository abstaction for note
type NoteRepository interface {
	CreateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (string, error)
	UpdateNote(ctx context.Context, noteID string, payload *entity.CreateUpdateNotePayload) (string, error)
	DeleteNote(ctx context.Context, noteID string) error
	GetNote(ctx context.Context, noteID string) (*entity.Note, error)
	GetNoteList(ctx context.Context, params *entity.GetNoteListParams) ([]*entity.Note, int64, error)
}
