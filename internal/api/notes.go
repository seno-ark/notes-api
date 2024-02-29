package api

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/entity"
	appErr "notes-api/pkg/error"
	"notes-api/pkg/utils"

	"github.com/go-chi/render"
)

func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := utils.NewResponse()

	payload := new(entity.Note)
	err := json.NewDecoder(r.Body).Decode(&payload)
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

	book, err := h.usecase.CreateNote(ctx, payload)
	if err != nil {
		status, message := appErr.ErrStatusCode(err)

		render.Status(r, status)
		render.JSON(w, r, resp.Set(message, nil))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp.Set("success", book))
}

func (h *handler) UpdateNote(w http.ResponseWriter, r *http.Request) {}

func (h *handler) DeleteNote(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetNote(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetNoteList(w http.ResponseWriter, r *http.Request) {}
