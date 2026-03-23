package handlers

import (
	"encoding/json"
	"net/http"
	"prompt-control-go/internal/models/prompt"
	"prompt-control-go/internal/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type PromptHandler struct {}

func (h *PromptHandler) Generate(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"create_prompt": services.Generation(chi.URLParam(r, "query")),
	})
}

func (h *PromptHandler) Refine(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"refine_prompt": services.SendRefinePrompt(chi.URLParam(r, "query")),
	})
}

func (h *PromptHandler) Enrich(w http.ResponseWriter, r *http.Request) {
	var rq models.EnrichRequest
	json.NewDecoder(r.Body).Decode(&rq)

	render.JSON(w, r, services.SendEnrichPrompt(rq))
}