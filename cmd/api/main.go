package main

import (
	"notes-api/internal/api"
	"notes-api/internal/repository"
	"notes-api/internal/usecase"
)

func main() {
	noteRepository := repository.NewNoteRepository()
	noteUsecase := usecase.NewNoteUsecase(noteRepository)
	handler := api.NewHandler(noteUsecase)
	handler.Routes()
}
