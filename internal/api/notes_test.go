package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"notes-api/internal/entity"
	"notes-api/internal/mocks"
	appErr "notes-api/pkg/error"
	"notes-api/pkg/utils"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	noteID, _ := utils.ULID()
	noteTitle := "Any Note"
	noteContent := "Just random new Note"

	testCases := []struct {
		Name               string
		Payload            *entity.CreateUpdateNotePayload
		MockResult         *entity.Note
		MockError          error
		ExpectedStatusCode int
	}{
		{
			Name: "Create Note Success",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			MockResult: &entity.Note{
				ID:      noteID,
				Title:   noteTitle,
				Content: noteContent,
			},
			ExpectedStatusCode: http.StatusCreated,
		},
		{
			Name: "Create Note Failed 400",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   "",
				Content: noteContent,
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "Create Note Failed 500",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			MockError:          appErr.NewErrInternalServer("failed to create note"),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tc.Payload)
			req, err := http.NewRequest("POST", "/notes", bytes.NewReader(jsonPayload))
			if err != nil {
				t.Fatal(err)
			}
			rec := httptest.NewRecorder()

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			if tc.MockResult != nil || tc.MockError != nil {
				mockNoteUsecase.On("CreateNote", context.Background(), tc.Payload).Return(tc.MockResult, tc.MockError).Once()
			}
			testHandler := NewHandler(mockNoteUsecase)
			testHandler.CreateNote(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Result().StatusCode)
		})
	}
}

func TestUpdateNote(t *testing.T) {
	noteID, _ := utils.ULID()
	noteTitle := "Any Note"
	noteContent := "Just another random Note"

	testCases := []struct {
		Name               string
		Payload            *entity.CreateUpdateNotePayload
		MockResult         *entity.Note
		MockError          error
		ExpectedStatusCode int
	}{
		{
			Name: "Update Note Success",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			MockResult: &entity.Note{
				ID:      noteID,
				Title:   noteTitle,
				Content: noteContent,
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name: "Update Note Failed 400",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   "",
				Content: noteContent,
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "Update Note Failed 500",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			MockError:          appErr.NewErrInternalServer("failed to update note"),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("note_id", noteID)
			jsonPayload, _ := json.Marshal(tc.Payload)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("PUT", "/notes/{note_id}", bytes.NewReader(jsonPayload))
			if err != nil {
				t.Fatal(err)
			}
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			if tc.MockResult != nil || tc.MockError != nil {
				mockNoteUsecase.On("UpdateNote", req.Context(), noteID, tc.Payload).Return(tc.MockResult, tc.MockError).Once()
			}
			testHandler := NewHandler(mockNoteUsecase)
			testHandler.UpdateNote(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Result().StatusCode)
		})
	}
}

func TestDeleteNote(t *testing.T) {
	noteID, _ := utils.ULID()

	testCases := []struct {
		Name               string
		NoteID             string
		MockError          error
		ExpectedStatusCode int
	}{
		{
			Name:               "Delete Note Success",
			NoteID:             noteID,
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Delete Note Failed 404",
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Name:               "Delete Note Failed 500",
			NoteID:             noteID,
			MockError:          appErr.NewErrInternalServer("failed to delete note"),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("note_id", tc.NoteID)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("DELETE", "/notes/{note_id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			if len(tc.NoteID) > 0 {
				mockNoteUsecase.On("DeleteNote", req.Context(), noteID).Return(tc.MockError).Once()
			}
			testHandler := NewHandler(mockNoteUsecase)
			testHandler.DeleteNote(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Result().StatusCode)
		})
	}
}
