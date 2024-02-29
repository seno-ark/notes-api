package repository

import (
	"context"
	"notes-api/internal/entity"
	appErr "notes-api/pkg/error"
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

		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(noteID))
	})

	t.Run("Success Get Created Note", func(t *testing.T) {
		note, err := testRepository.GetNote(context.Background(), noteID)

		assert.NoError(t, err)
		assert.NotEmpty(t, note)

		assert.Equal(t, noteID, note.ID)
		assert.Equal(t, "New Note v1", note.Title)
		assert.Equal(t, "TestNotes v1", note.Content)
		assert.Equal(t, note.CreatedAt, note.UpdatedAt)
	})

	t.Run("Success Update Note", func(t *testing.T) {
		payload := &entity.Note{
			ID:      noteID,
			Title:   "Updated Note v2",
			Content: "TestNotes v2",
		}
		updatedNoteID, err := testRepository.UpdateNote(context.Background(), payload)

		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(updatedNoteID))
		assert.Equal(t, noteID, updatedNoteID)
	})

	t.Run("Success Get Updated Note", func(t *testing.T) {
		note, err := testRepository.GetNote(context.Background(), noteID)

		assert.NoError(t, err)
		assert.NotEmpty(t, note)

		assert.Equal(t, noteID, note.ID)
		assert.Equal(t, "Updated Note v2", note.Title)
		assert.Equal(t, "TestNotes v2", note.Content)
		assert.NotEqual(t, note.CreatedAt, note.UpdatedAt)
	})

	t.Run("Success Get Note List", func(t *testing.T) {
		params := &entity.GetNoteListParams{
			Offset: 0,
			Limit:  10,
			Sort:   "-created_at",
			Search: "",
		}
		notes, total, err := testRepository.GetNoteList(context.Background(), params)

		assert.NoError(t, err)
		assert.NotEmpty(t, total)
		assert.NotEmpty(t, notes)
	})

	t.Run("Search Get Note List Found", func(t *testing.T) {
		params := &entity.GetNoteListParams{
			Offset: 0,
			Limit:  10,
			Sort:   "-created_at",
			Search: "v2",
		}
		notes, total, err := testRepository.GetNoteList(context.Background(), params)

		assert.NoError(t, err)
		assert.NotEmpty(t, total)
		assert.NotEmpty(t, notes)
	})

	t.Run("Search Get Note List Not Found", func(t *testing.T) {
		params := &entity.GetNoteListParams{
			Offset: 0,
			Limit:  10,
			Sort:   "-created_at",
			Search: "3c94n823m209n",
		}
		notes, total, err := testRepository.GetNoteList(context.Background(), params)

		assert.NoError(t, err)
		assert.Empty(t, total)
		assert.Empty(t, notes)
	})

	t.Run("Success Delete Note", func(t *testing.T) {
		err := testRepository.DeleteNote(context.Background(), noteID)

		assert.NoError(t, err)
	})

	t.Run("Get Deleted Note", func(t *testing.T) {
		note, err := testRepository.GetNote(context.Background(), noteID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrNotFound)
		assert.Empty(t, note)
	})
}
