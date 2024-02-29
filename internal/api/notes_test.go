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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	nowUtc := time.Now().UTC()
	noteID, _ := utils.ULID()
	noteTitle := "Any Note"
	noteContent := "Just random new Note"

	testCases := []struct {
		Name       string
		Payload    *entity.CreateUpdateNotePayload
		Result     *entity.Note
		Error      error
		StatusCode int
	}{
		{
			Name: "Create Note Success",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			Result: &entity.Note{
				ID:        noteID,
				Title:     noteTitle,
				Content:   noteContent,
				CreatedAt: nowUtc,
				UpdatedAt: nowUtc,
			},
			StatusCode: http.StatusCreated,
		},
		{
			Name: "Create Note Failed 400",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   "",
				Content: noteContent,
			},
			StatusCode: http.StatusBadRequest,
		},
		{
			Name: "Create Note Failed 500",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			Error:      appErr.NewErrInternalServer("failed to create note"),
			StatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockNoteUsecase := mocks.NewNoteUsecase(t)

			if tc.Result != nil || tc.Error != nil {
				mockNoteUsecase.On("CreateNote", context.Background(), tc.Payload).Return(tc.Result, tc.Error).Once()
			}

			testHandler := NewHandler(mockNoteUsecase)

			rec := httptest.NewRecorder()

			jsonPayload, _ := json.Marshal(tc.Payload)
			req, err := http.NewRequest("POST", "/notes", bytes.NewReader(jsonPayload))
			if err != nil {
				t.Fatal(err)
			}

			testHandler.CreateNote(rec, req)

			assert.Equal(t, tc.StatusCode, rec.Result().StatusCode)
		})
	}
}
