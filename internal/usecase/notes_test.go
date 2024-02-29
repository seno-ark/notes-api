package usecase

import (
	"context"
	"notes-api/internal/entity"
	"notes-api/internal/mocks"
	appErr "notes-api/pkg/error"
	"notes-api/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	ctx := context.Background()

	payload := &entity.Note{
		Title:   "Test Create Note",
		Content: "Created from TestCreateNote",
	}

	t.Run("Success Create Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		id, _ := utils.ULID()
		nowUtc := time.Now().UTC()
		result := &entity.Note{
			ID:        id,
			Title:     payload.Title,
			Content:   payload.Content,
			CreatedAt: nowUtc,
			UpdatedAt: nowUtc,
		}

		mockNoteRepository.On("CreateNote", ctx, payload).Return(id, nil).Once()
		mockNoteRepository.On("GetNote", ctx, id).Return(result, nil).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		createdNote, err := testNoteUsecase.CreateNote(ctx, payload)

		assert.NoError(t, err)
		assert.Equal(t, result, createdNote)
	})

	t.Run("Failed Create Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		errCreateNote := appErr.NewErrInternalServer("failed to create note")
		mockNoteRepository.On("CreateNote", ctx, payload).Return("", errCreateNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		createdNote, err := testNoteUsecase.CreateNote(ctx, payload)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrInternalServer)
		assert.Empty(t, createdNote)
	})

}

func TestUpdateNote(t *testing.T) {
	ctx := context.Background()

	id, _ := utils.ULID()
	payload := &entity.Note{
		ID:      id,
		Title:   "Test Update Note",
		Content: "Updated from TestUpdateNote",
	}

	nowUtc := time.Now().UTC()
	expectedResult := &entity.Note{
		ID:        id,
		Title:     payload.Title,
		Content:   payload.Content,
		CreatedAt: nowUtc,
		UpdatedAt: nowUtc,
	}

	t.Run("Success Update Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		mockNoteRepository.On("GetNote", ctx, id).Return(expectedResult, nil).Once()
		mockNoteRepository.On("UpdateNote", ctx, payload).Return(id, nil).Once()
		mockNoteRepository.On("GetNote", ctx, id).Return(expectedResult, nil).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		updatedNote, err := testNoteUsecase.UpdateNote(ctx, payload)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, updatedNote)
	})

	t.Run("Failed Update Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		mockNoteRepository.On("GetNote", ctx, id).Return(expectedResult, nil).Once()

		errUpdateNote := appErr.NewErrInternalServer("failed to update note")
		mockNoteRepository.On("UpdateNote", ctx, payload).Return("", errUpdateNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		createdNote, err := testNoteUsecase.UpdateNote(ctx, payload)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrInternalServer)
		assert.Empty(t, createdNote)
	})

	t.Run("Update Non Existed Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		errGetNote := appErr.NewErrNotFound("note not found")
		mockNoteRepository.On("GetNote", ctx, id).Return(nil, errGetNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		createdNote, err := testNoteUsecase.UpdateNote(ctx, payload)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrNotFound)
		assert.Empty(t, createdNote)
	})
}

func TestDeleteNote(t *testing.T) {
	ctx := context.Background()

	id, _ := utils.ULID()
	nowUtc := time.Now().UTC()
	expectedResult := &entity.Note{
		ID:        id,
		Title:     "Test Delete Note",
		Content:   "Deleted from TestDeleteNote",
		CreatedAt: nowUtc,
		UpdatedAt: nowUtc,
	}

	t.Run("Success Delete Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		mockNoteRepository.On("GetNote", ctx, id).Return(expectedResult, nil).Once()
		mockNoteRepository.On("DeleteNote", ctx, id).Return(nil).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		err := testNoteUsecase.DeleteNote(ctx, id)

		assert.NoError(t, err)
	})

	t.Run("Failed Delete Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		mockNoteRepository.On("GetNote", ctx, id).Return(expectedResult, nil).Once()

		errDeleteNote := appErr.NewErrInternalServer("failed to delete note")
		mockNoteRepository.On("DeleteNote", ctx, id).Return(errDeleteNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		err := testNoteUsecase.DeleteNote(ctx, id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrInternalServer)
	})

	t.Run("Delete Non Existed Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		errGetNote := appErr.NewErrNotFound("note not found")
		mockNoteRepository.On("GetNote", ctx, id).Return(nil, errGetNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		err := testNoteUsecase.DeleteNote(ctx, id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrNotFound)
	})
}

func TestGetNote(t *testing.T) {
	ctx := context.Background()

	id, _ := utils.ULID()
	nowUtc := time.Now().UTC()
	expectedResult := &entity.Note{
		ID:        id,
		Title:     "Test Delete Note",
		Content:   "Deleted from TestDeleteNote",
		CreatedAt: nowUtc,
		UpdatedAt: nowUtc,
	}

	t.Run("Success Get Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		mockNoteRepository.On("GetNote", ctx, id).Return(expectedResult, nil).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		note, err := testNoteUsecase.GetNote(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, note)
	})

	t.Run("Failed Get Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		errGetNote := appErr.NewErrInternalServer("failed to get note")
		mockNoteRepository.On("GetNote", ctx, id).Return(nil, errGetNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		note, err := testNoteUsecase.GetNote(ctx, id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrInternalServer)
		assert.Empty(t, note)
	})

	t.Run("Get Non Existed Note", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		errGetNote := appErr.NewErrNotFound("note not found")
		mockNoteRepository.On("GetNote", ctx, id).Return(nil, errGetNote).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		note, err := testNoteUsecase.GetNote(ctx, id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrNotFound)
		assert.Empty(t, note)
	})
}

func TestGetNoteList(t *testing.T) {
	ctx := context.Background()

	id, _ := utils.ULID()
	nowUtc := time.Now().UTC()
	expectedResults := []*entity.Note{
		{
			ID:        id,
			Title:     "Test Delete Note",
			Content:   "Deleted from TestDeleteNote",
			CreatedAt: nowUtc,
			UpdatedAt: nowUtc,
		},
	}

	params := &entity.GetNoteListParams{}

	t.Run("Success Get Notes", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		expectedTotal := int64(1)
		mockNoteRepository.On("GetNoteList", ctx, params).Return(expectedResults, expectedTotal, nil).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		notes, total, err := testNoteUsecase.GetNoteList(ctx, params)

		assert.NoError(t, err)
		assert.Equal(t, expectedResults, notes)
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("Success Get Zero Notes", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		mockNoteRepository.On("GetNoteList", ctx, params).Return(nil, int64(0), nil).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		notes, total, err := testNoteUsecase.GetNoteList(ctx, params)

		assert.NoError(t, err)
		assert.Empty(t, notes)
		assert.Empty(t, total)
	})

	t.Run("Failed Get Notes", func(t *testing.T) {
		mockNoteRepository := mocks.NewNoteRepository(t)

		params := &entity.GetNoteListParams{}

		errGetNotes := appErr.NewErrInternalServer("failed to get notes")
		mockNoteRepository.On("GetNoteList", ctx, params).Return(nil, int64(0), errGetNotes).Once()

		testNoteUsecase := NewNoteUsecase(mockNoteRepository)
		notes, total, err := testNoteUsecase.GetNoteList(ctx, params)

		assert.Error(t, err)
		assert.ErrorIs(t, err, appErr.ErrInternalServer)
		assert.Empty(t, notes)
		assert.Empty(t, total)
	})
}
