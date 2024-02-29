package api

import (
	"notes-api/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	usecase  internal.NoteUsecase
	validate *validator.Validate
}

func NewHandler(usecase internal.NoteUsecase) *handler {
	return &handler{
		usecase:  usecase,
		validate: validator.New(),
	}
}

func Routes(r *chi.Mux, h *handler) {
	r.Post("/notes", h.CreateNote)
	r.Put("/notes/{note_id}", h.UpdateNote)
	r.Delete("/notes/{note_id}", h.DeleteNote)
	r.Get("/notes/{note_id}", h.GetNote)
	r.Get("/notes", h.GetNoteList)
}
