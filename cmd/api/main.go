package main

import (
	"notes-api/config"
	"notes-api/internal/api"
	"notes-api/internal/repository"
	"notes-api/internal/usecase"
	"notes-api/pkg/database"
)

func main() {
	conf := config.GetConfig()

	db, err := database.Postgres(conf)
	if err != nil {
		panic(err)
	}

	noteRepository := repository.NewNoteRepository(db)
	noteUsecase := usecase.NewNoteUsecase(noteRepository)
	handler := api.NewHandler(noteUsecase)
	handler.Routes()
}
