package model

import "github.com/google/uuid"

type CreatePreloved struct {
	ID_Preloved uuid.UUID	`json:"-"`
	ID_User       uuid.UUID `json:"-"`
	Title       string `json:"title" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Price       int   `json:"price,string" binding:"required"`
	Condition   string `json:"condition" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdatePreloved struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	Price       int   `json:"price,string"`
	Condition   string `json:"condition"`
	Description string `json:"description"`
}