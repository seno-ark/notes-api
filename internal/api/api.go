package api

import "notes-api/internal"

type handler struct {
	usecase internal.NoteUsecase
}

func NewHandler(usecase internal.NoteUsecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) Routes() {}
