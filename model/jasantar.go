package model

import "github.com/google/uuid"

type CreateJasantar struct {
	ID_Jasantar uuid.UUID `json:"-"`
	ID_User     uuid.UUID `json:"-"`
	Title       string    `json:"title" binding:"required"`
	Area        string    `json:"area" binding:"required"`
	Category    string    `json:"category" binding:"required"`
}

type UpdateJasantar struct {
	ID_Jasantar uuid.UUID `json:"-"`
	ID_User     uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Area        string    `json:"area"`
	Category    string    `json:"category"`
}