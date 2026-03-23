package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=password dbname=mydb port=5432 sslmode=disable"
    
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Произошла ошибка при коннекте к DB")
	}
}