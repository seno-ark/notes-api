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
			Name: "Create Note 201",
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

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/notes", bytes.NewReader(jsonPayload))
			if err != nil {
				t.Fatal(err)
			}

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
		NoteID             string
		Payload            *entity.CreateUpdateNotePayload
		MockResult         *entity.Note
		MockError          error
		ExpectedStatusCode int
	}{
		{
			Name:   "Update Note 200",
			NoteID: noteID,
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
			Name:   "Update Note Failed 400",
			NoteID: noteID,
			Payload: &entity.CreateUpdateNotePayload{
				Title:   "",
				Content: noteContent,
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name:   "Update Note Failed 404",
			NoteID: "XXX",
			Payload: &entity.CreateUpdateNotePayload{
				Title:   noteTitle,
				Content: noteContent,
			},
			MockError:          appErr.NewErrNotFound("note not found"),
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Name:   "Update Note Failed 500",
			NoteID: noteID,
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
			jsonPayload, _ := json.Marshal(tc.Payload)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("PUT", "/notes/{note_id}", bytes.NewReader(jsonPayload))
			if err != nil {
				t.Fatal(err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("note_id", tc.NoteID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			if tc.MockResult != nil || tc.MockError != nil {
				mockNoteUsecase.On("UpdateNote", req.Context(), tc.NoteID, tc.Payload).Return(tc.MockResult, tc.MockError).Once()
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
			Name:               "Delete Note 200",
			NoteID:             noteID,
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Delete Note Failed 404",
			NoteID:             "XXX",
			MockError:          appErr.NewErrNotFound("note not found"),
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
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("DELETE", "/notes/{note_id}", nil)
			if err != nil {
				t.Fatal(err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("note_id", tc.NoteID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			mockNoteUsecase.On("DeleteNote", req.Context(), tc.NoteID).Return(tc.MockError).Once()
			testHandler := NewHandler(mockNoteUsecase)
			testHandler.DeleteNote(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Result().StatusCode)
		})
	}
}

func TestGetNote(t *testing.T) {
	noteID, _ := utils.ULID()
	noteTitle := "Any Note"
	noteContent := "Just another random Note"

	testCases := []struct {
		Name               string
		NoteID             string
		MockResult         *entity.Note
		MockError          error
		ExpectedStatusCode int
	}{
		{
			Name:   "Get Note 200",
			NoteID: noteID,
			MockResult: &entity.Note{
				ID:      noteID,
				Title:   noteTitle,
				Content: noteContent,
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Get Note Failed 404",
			NoteID:             "XXX",
			MockError:          appErr.NewErrNotFound("note not found"),
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Name:               "Get Note Failed 500",
			NoteID:             noteID,
			MockError:          appErr.NewErrInternalServer("failed to get note"),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("DELETE", "/notes/{note_id}", nil)
			if err != nil {
				t.Fatal(err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("note_id", tc.NoteID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			mockNoteUsecase.On("GetNote", req.Context(), tc.NoteID).Return(tc.MockResult, tc.MockError).Once()
			testHandler := NewHandler(mockNoteUsecase)
			testHandler.GetNote(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Result().StatusCode)
		})
	}
}

type getNoteListQueryParams struct {
	Page   string
	Count  string
	Sort   string
	Search string
}

func (p getNoteListQueryParams) ToUsecaseParam() *entity.GetNoteListParams {
	page, count := utils.Pagination(p.Page, p.Count)

	return &entity.GetNoteListParams{
		Offset: (page - 1) * count,
		Limit:  count,
		Sort:   p.Sort,
		Search: p.Search,
	}
}

func TestGetNoteList(t *testing.T) {
	noteID, _ := utils.ULID()
	noteTitle := "Any Note"
	noteContent := "Just another random Note"

	testCases := []struct {
		Name               string
		Params             getNoteListQueryParams
		MockResults        []*entity.Note
		MockTotalResult    int64
		MockError          error
		ExpectedStatusCode int
	}{
		{
			Name: "Get Note List 200",
			Params: getNoteListQueryParams{
				Page:  "1",
				Count: "10",
			},
			MockResults: []*entity.Note{
				{
					ID:      noteID,
					Title:   noteTitle,
					Content: noteContent,
				},
			},
			MockTotalResult:    1,
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name: "Get Empty Note List 200",
			Params: getNoteListQueryParams{
				Page:   "1",
				Count:  "10",
				Search: "23rfwsvszv",
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name: "Get Note List Failed 500",
			Params: getNoteListQueryParams{
				Page:  "1",
				Count: "10",
			},
			MockError:          appErr.NewErrInternalServer("failed to get note list"),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/notes", nil)
			if err != nil {
				t.Fatal(err)
			}

			q := req.URL.Query()
			q.Add("page", tc.Params.Page)
			q.Add("count", tc.Params.Count)
			q.Add("sort", tc.Params.Sort)
			q.Add("search", tc.Params.Search)
			req.URL.RawQuery = q.Encode()

			mockNoteUsecase := mocks.NewNoteUsecase(t)
			mockNoteUsecase.
				On("GetNoteList", req.Context(), tc.Params.ToUsecaseParam()).
				Return(tc.MockResults, tc.MockTotalResult, tc.MockError).
				Once()
			testHandler := NewHandler(mockNoteUsecase)
			testHandler.GetNoteList(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Result().StatusCode)
		})
	}
}
