package repository

import (
	"notes-api/internal/entity"
	"time"
)

// Note mapping external data to database column or vice versa
type Note struct {
	ID        string    `gorm:"column:id;primary_key"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// FromDto convert entity.CreateUpdateNotePayload to repository.Note model
func (m *Note) FromDto(e *entity.CreateUpdateNotePayload) {
	m.ID = e.ID
	m.Title = e.Title
	m.Content = e.Content
}

// ToEntity convert repository.Note model to entity.CreateUpdateNotePayload
func (m *Note) ToEntity() *entity.Note {
	return &entity.Note{
		ID:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
