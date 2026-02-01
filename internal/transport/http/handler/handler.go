package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"thought-RestAPI/internal/domain"
	"thought-RestAPI/internal/dto"
	"thought-RestAPI/internal/usecase"
)

type ThoughtHandler struct {
	uc *usecase.Thought
}

func NewHandlers(uc *usecase.Thought) *ThoughtHandler {
	return &ThoughtHandler{
		uc: uc,
	}
}

func (h *ThoughtHandler) GetRandomThought(w http.ResponseWriter, r *http.Request) {
	thought, err := h.uc.GetRandomThought(r.Context())
	if err != nil {
		if errors.Is(err, domain.ErrThoughtNotFound) {
			writeError(w, r, http.StatusNotFound, "thought not found")
			return
		}
		writeError(w, r, http.StatusInternalServerError, "internal error")
		return
	}

	render.JSON(w, r, thought)
}

func (h *ThoughtHandler) CreateThought(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.CreateThoughtRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Text == "" || req.Author == "" {
		writeError(w, r, http.StatusBadRequest, "text and author are required")
		return
	}

	id, err := h.uc.CreateThought(
		r.Context(),
		usecase.CreateThoughtInput{
			Text:   req.Text,
			Author: req.Author,
		},
	)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, dto.CreateThoughtResponse{ID: id})

}

func (h *ThoughtHandler) UpdateThought(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, r, http.StatusBadRequest, "invalid id")
		return
	}

	defer r.Body.Close()

	var req dto.UpdateThoughtRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Text == "" || req.Author == "" {
		writeError(w, r, http.StatusBadRequest, "text and author are required")
		return
	}

	err = h.uc.UpdateThought(r.Context(), usecase.UpdateThoughtInput{
		ID:     id,
		Text:   req.Text,
		Author: req.Author,
	})
	if err != nil {
		if errors.Is(err, domain.ErrThoughtNotFound) {
			writeError(w, r, http.StatusNotFound, "thought not found")
			return
		}
		writeError(w, r, http.StatusInternalServerError, "internal error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ThoughtHandler) DeleteThought(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, r, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.uc.DeleteThought(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrThoughtNotFound) {
			writeError(w, r, http.StatusNotFound, "thought not found")
			return
		}
		writeError(w, r, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	render.Status(r, code)
	render.JSON(w, r, map[string]string{"error": msg})
}
