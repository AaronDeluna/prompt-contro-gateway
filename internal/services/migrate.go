package services

import (
	"log"
	"prompt-control-go/internal/db"

	"github.com/pressly/goose/v3"
)

func UpMigration() {
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		log.Fatal(err)
	}
	log.Println("Миграция успешно прошла")
}