// Package main provides main package for api application
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"notes-api/config"
	"notes-api/internal/api"
	"notes-api/internal/repository"
	"notes-api/internal/usecase"
	"notes-api/pkg/database"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "notes-api/cmd/api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Simple Notes API
// @version 1.0
// @description This is a simple notes api server

// @host localhost:9000
// @BasePath /
func main() {
	conf := config.GetConfig("./")

	db, err := database.Postgres(conf)
	if err != nil {
		panic(err)
	}

	noteRepository := repository.NewNoteRepository(db)
	noteUsecase := usecase.NewNoteUsecase(noteRepository)
	handler := api.NewHandler(noteUsecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))

	api.Routes(r, handler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	slog.Info("start http server", slog.Any("port", conf.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), r)
	if err != nil {
		slog.Error("error start http server", "port", conf.Port, "err", err)
		panic(err)
	}
}
