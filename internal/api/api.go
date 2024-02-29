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
	r.Post("/books", h.CreateNote)
	r.Put("/books/{book_id}", h.UpdateNote)
	r.Delete("/books/{book_id}", h.DeleteNote)
	r.Get("/books/{book_id}", h.GetNote)
	r.Get("/books", h.GetNoteList)
}
