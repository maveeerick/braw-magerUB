package model

import "github.com/google/uuid"

type CreateKomunitasbraw struct {
	IdKomunitasbraw uuid.UUID `json:"-"`
	IdUser          uuid.UUID `json:"idUser" binding:"required"`
	Title           string    `json:"title" binding:"required"`
	Category        string    `json:"category" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	LinkWebsite     string    `json:"linkwebsite" binding:"required"`
}

type UpdateKomunitasbraw struct {
	LinkWebsite string `json:"linkwebsite"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
