package api

import (
	"notes-api/internal"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	usecase internal.NoteUsecase
}

func NewHandler(usecase internal.NoteUsecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func Routes(r *chi.Mux, h *handler) {
	r.Post("/books", h.CreateNote)
	r.Put("/books/{book_id}", h.UpdateNote)
	r.Delete("/books/{book_id}", h.DeleteNote)
	r.Get("/books/{book_id}", h.GetNote)
	r.Get("/books", h.GetNoteList)
}
