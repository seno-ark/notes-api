package entity

import "time"

type Note struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetNoteListFilter struct {
	Offset int
	Limit  int
	Sort   string
	Search string
}
