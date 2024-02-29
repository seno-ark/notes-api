// Package api provides application routes and handler for api
package api

import (
	"notes-api/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// Handler provides all endpoint handlers
type Handler struct {
	usecase  internal.NoteUsecase
	validate *validator.Validate
}

// NewHandler returns new *Handler
func NewHandler(usecase internal.NoteUsecase) *Handler {
	return &Handler{
		usecase:  usecase,
		validate: validator.New(),
	}
}

// Routes map routes and handlers
func Routes(r *chi.Mux, h *Handler) {
	r.Post("/notes", h.CreateNote)
	r.Put("/notes/{note_id}", h.UpdateNote)
	r.Delete("/notes/{note_id}", h.DeleteNote)
	r.Get("/notes/{note_id}", h.GetNote)
	r.Get("/notes", h.GetNoteList)
}
