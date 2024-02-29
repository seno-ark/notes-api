package api

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/entity"
	appErr "notes-api/pkg/error"
	"notes-api/pkg/utils"

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

func (h *handler) UpdateNote(w http.ResponseWriter, r *http.Request) {}

func (h *handler) DeleteNote(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetNote(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetNoteList(w http.ResponseWriter, r *http.Request) {}
