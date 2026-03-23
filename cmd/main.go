package main

import (
	"log"
	"net/http"
	"prompt-control-go/internal/db"
	"prompt-control-go/internal/handlers"
	"prompt-control-go/internal/services"

	"github.com/go-chi/chi"
)

const (
	SERVER_PORT = ":8089"
)

func main() {
	initDb()

	promptHandler := &handlers.PromptHandler{}

	r := chi.NewRouter()
	r.Route("/prompts", func(r chi.Router) {
		r.Post("/generate/{query}", promptHandler.Generate)
		r.Post("/refine/{query}", promptHandler.Refine)
		r.Post("/enrich", promptHandler.Enrich)
	})

	log.Println("Старт сервера на порту: ", SERVER_PORT)
	log.Fatal(http.ListenAndServe(SERVER_PORT, r))
}

func initDb() {
	db.Connect()
	services.UpMigration()
}
