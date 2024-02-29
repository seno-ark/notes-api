package repository

import (
	"notes-api/internal/entity"
	"time"
)

type Note struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Note) FromDto(e *entity.CreateUpdateNotePayload) {
	m.ID = e.ID
	m.Title = e.Title
	m.Content = e.Content
}

func (m *Note) ToEntity() *entity.Note {
	return &entity.Note{
		ID:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
