package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"notes-api/internal/entity"
	"notes-api/internal/mocks"
	"notes-api/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	id, _ := utils.ULID()
	nowUtc := time.Now().UTC()

	payload := &entity.Note{
		Title:   "New Note",
		Content: "Just random new Note",
	}

	testCases := []struct {
		Name       string
		Payload    *entity.Note
		Result     *entity.Note
		Error      error
		StatusCode int
	}{
		{
			Name:    "Create Note Success",
			Payload: payload,
			Result: &entity.Note{
				ID:        id,
				Title:     payload.Title,
				Content:   payload.Content,
				CreatedAt: nowUtc,
				UpdatedAt: nowUtc,
			},
			Error:      nil,
			StatusCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockNoteUsecase := mocks.NewNoteUsecase(t)
			mockNoteUsecase.On("CreateBook", context.Background(), tc.Payload).Return(tc.Result, tc.Error).Once()

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
