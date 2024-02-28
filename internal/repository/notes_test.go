package repository

import (
	"context"
	"notes-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotes(t *testing.T) {

	var (
		noteID string
		err    error
	)

	t.Run("Success Create Note", func(t *testing.T) {
		payload := &entity.Note{
			Title:   "New Note v1",
			Content: "TestNotes v1",
		}
		noteID, err = testRepository.CreateNote(context.Background(), payload)

		assert.Empty(t, err)
		assert.NotEqual(t, 0, len(noteID))
	})

	t.Run("Success Update Note", func(t *testing.T) {
		payload := &entity.Note{
			ID:      noteID,
			Title:   "Updated Note v2",
			Content: "TestNotes v2",
		}
		updatedNoteID, err := testRepository.UpdateNote(context.Background(), payload)

		assert.Empty(t, err)
		assert.NotEqual(t, 0, len(updatedNoteID))
		assert.Equal(t, noteID, updatedNoteID)
	})

	t.Run("Success Get Note", func(t *testing.T) {
		note, err := testRepository.GetNote(context.Background(), noteID)

		assert.Empty(t, err)
		assert.NotEmpty(t, note)

		assert.Equal(t, "Updated Note v2", note.Title)
		assert.Equal(t, "TestNotes v2", note.Content)
	})

	t.Run("Success Get Note List", func(t *testing.T) {
		filter := &entity.GetNoteListFilter{
			Offset: 0,
			Limit:  10,
			Sort:   "id",
			Search: "",
		}
		notes, total, err := testRepository.GetNoteList(context.Background(), filter)

		assert.Empty(t, err)
		assert.NotEmpty(t, total)
		assert.NotEmpty(t, notes)
	})
}
