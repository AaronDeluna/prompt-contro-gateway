package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type Handler struct {

}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
    render.JSON(w, r, map[string]string{"message": "pong"})
}