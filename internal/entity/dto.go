package entity

// CreateUpdateNotePayload is payload used for creating or updating note
type CreateUpdateNotePayload struct {
	ID      string `json:"-"`
	Title   string `json:"title" validate:"required,min=1,max=255" example:"To Do list"`
	Content string `json:"content" validate:"required,min=1,max=1000" example:"1. Nothing."`
}

// GetNoteListParams is parameter used for filtering and sorting to get list of notes
type GetNoteListParams struct {
	Offset int
	Limit  int
	Sort   string
	Search string
}
