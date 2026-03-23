package models

type Prompt struct {
	Id uint `gorm:"primaryKey"`
	Name string
	Subject string
}

type EnrichRequest struct {
    Prompt string   `json:"prompt"`
    Additions    []string `json:"additions"`
}

type EnrichResponse struct {
	Name string `json:"name"`
	Prompt string `json:"prompt"`
}