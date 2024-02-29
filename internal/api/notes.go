package api

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/entity"
	appErr "notes-api/pkg/error"
	"notes-api/pkg/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CreateNote
// @Summary			Create Note.
// @Description		Create new Note.
// @Tags			Notes
// @Accept			json
// @Produce			json
// @Param 			note	body	entity.CreateUpdateNotePayload	true	"note data"
// @Success			201		{object}	utils.Response
// @Failure			400		{object}	utils.Response
// @Failure			401		{object}	utils.Response
// @Failure			422		{object}	utils.Response
// @Failure			500		{object}	utils.Response
// @Router	/notes [post]
func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := utils.NewResponse()

	payload := new(entity.CreateUpdateNotePayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil))
		return
	}

	err = h.validate.Struct(payload)
	if err != nil {
		errs := utils.ParseValidatorErr(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil).AddErrValidation(errs))
		return
	}

	note, err := h.usecase.CreateNote(ctx, payload)
	if err != nil {
		status, message := appErr.ErrStatusCode(err)

		render.Status(r, status)
		render.JSON(w, r, resp.Set(message, nil))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp.Set("success", note))
}

// UpdateNote
// @Summary			Update Note.
// @Description		Update existing Note.
// @Tags			Notes
// @Accept			json
// @Produce			json
// @Param 			note_id	path	string							true	"Note ID" example(01HQSH92SNYQVCBDSD38XNBRYM)
// @Param 			note	body	entity.CreateUpdateNotePayload	true	"note data"
// @Success			200		{object}	utils.Response
// @Failure			400		{object}	utils.Response
// @Failure			401		{object}	utils.Response
// @Failure			422		{object}	utils.Response
// @Failure			500		{object}	utils.Response
// @Router	/notes/{note_id} [put]
func (h *handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := utils.NewResponse()

	noteID := chi.URLParam(r, "note_id")
	if noteID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("note not found", nil))
		return
	}

	payload := new(entity.CreateUpdateNotePayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil))
		return
	}

	err = h.validate.Struct(payload)
	if err != nil {
		errs := utils.ParseValidatorErr(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil).AddErrValidation(errs))
		return
	}

	note, err := h.usecase.UpdateNote(ctx, noteID, payload)
	if err != nil {
		status, message := appErr.ErrStatusCode(err)

		render.Status(r, status)
		render.JSON(w, r, resp.Set(message, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", note))
}

// DeleteNote
// @Summary			Delete note.
// @Description		Delete note.
// @Tags			Notes
// @Produce			json
// @Param			note_id			path			string	 true	"Note ID" example(01HQSH92SNYQVCBDSD38XNBRYM)
// @Success			200 			{object}		utils.Response
// @Failure			400				{object}		utils.Response
// @Failure			401				{object}		utils.Response
// @Failure			422				{object}		utils.Response
// @Failure			500				{object}		utils.Response
// @Router	/notes/{note_id} [delete]
func (h *handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := utils.NewResponse()

	noteID := chi.URLParam(r, "note_id")
	if noteID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("note not found", nil))
		return
	}

	err := h.usecase.DeleteNote(ctx, noteID)
	if err != nil {
		status, message := appErr.ErrStatusCode(err)

		render.Status(r, status)
		render.JSON(w, r, resp.Set(message, nil))
		return
	}

	result := map[string]any{
		"id": noteID,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", result))
}

// GetNote
// @Summary			Get note by note ID.
// @Description		Get note by note ID.
// @Tags			Notes
// @Produce			json
// @Param			note_id			path			string	 true	"Note ID"
// @Success			200 			{object}		utils.Response
// @Failure			400				{object}		utils.Response
// @Failure			401				{object}		utils.Response
// @Failure			422				{object}		utils.Response
// @Failure			500				{object}		utils.Response
// @Router	/notes/{note_id} [get]
func (h *handler) GetNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := utils.NewResponse()

	noteID := chi.URLParam(r, "note_id")
	if noteID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("note not found", nil))
		return
	}

	note, err := h.usecase.GetNote(ctx, noteID)
	if err != nil {
		status, message := appErr.ErrStatusCode(err)

		render.Status(r, status)
		render.JSON(w, r, resp.Set(message, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", note))
}

// GetNoteList
// @Summary			Get list of note.
// @Description		Get list of note.
// @Tags			Notes
// @Produce			json
// @Param			page			query			int	     false	"Pagination page number (default 1, max 500)"				example(1)
// @Param			count			query			int	     false	"Pagination data limit  (default 10, max 100)"				example(10)
// @Param			sort			query			string	 false	"Data sorting (value: id/title/created_at/updated_at). For desc order, use prefix '-'"	example(-created_at)
// @Param			search			query			string	 false	"Keyword for searching note by title or content" 			example(to do list)
// @Success			200 			{object}		utils.Response
// @Failure			400				{object}		utils.Response
// @Failure			401				{object}		utils.Response
// @Failure			422				{object}		utils.Response
// @Failure			500				{object}		utils.Response
// @Router	/notes [get]
func (h *handler) GetNoteList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := utils.NewResponse()
	q := r.URL.Query()

	page, count := utils.Pagination(q.Get("page"), q.Get("count"))

	params := &entity.GetNoteListParams{
		Offset: (page - 1) * count,
		Limit:  count,
		Sort:   q.Get("sort"),
		Search: q.Get("search"),
	}
	notes, total, err := h.usecase.GetNoteList(ctx, params)
	if err != nil {
		status, message := appErr.ErrStatusCode(err)

		render.Status(r, status)
		render.JSON(w, r, resp.Set(message, nil))
		return
	}

	resp.AddMeta(page, count, total)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", notes))
}
