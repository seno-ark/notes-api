package repository

import (
	"context"
	"notes-api/internal"
	"notes-api/internal/entity"
)

type noteRepository struct{}

func NewNoteRepository() internal.NoteRepository {
	return &noteRepository{}
}

func (r *noteRepository) CreateNote(ctx context.Context, payload *entity.Note) (string, error) {
	return "", nil
}

func (r *noteRepository) UpdateNote(ctx context.Context, payload *entity.Note) (string, error) {
	return "", nil
}

func (r *noteRepository) DeleteNote(ctx context.Context, bookID string) error {
	return nil
}

func (r *noteRepository) GetNote(ctx context.Context, bookID string) (*entity.Note, error) {
	return nil, nil
}

func (r *noteRepository) GetNoteList(ctx context.Context, filter *entity.GetNoteListFilter) ([]*entity.Note, int64, error) {
	return nil, 0, nil
}
