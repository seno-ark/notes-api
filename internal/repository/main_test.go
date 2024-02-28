package repository

import (
	"log"
	"notes-api/config"
	"notes-api/internal"
	"notes-api/pkg/database"
	"os"
	"testing"
)

var (
	testRepository internal.NoteRepository
)

func TestMain(m *testing.M) {
	var err error

	config := config.GetConfig()

	testDB, err := database.Postgres(config)
	if err != nil {
		log.Fatal("cannot connect postgres:", err)
	}

	testRepository = NewNoteRepository(testDB)

	os.Exit(m.Run())
}
