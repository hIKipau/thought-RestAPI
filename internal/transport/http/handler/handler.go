package handler

import (
	"net/http"
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

}
func (h *ThoughtHandler) CreateThought(w http.ResponseWriter, r *http.Request) {}
func (h *ThoughtHandler) UpdateThought(w http.ResponseWriter, r *http.Request) {}
func (h *ThoughtHandler) DeleteThought(w http.ResponseWriter, r *http.Request) {}
