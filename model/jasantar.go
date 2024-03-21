package model

import "github.com/google/uuid"

type CreateJasantar struct {
	IdJasantar uuid.UUID `json:"-"`
	IdUser     uuid.UUID `json:"idUser" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Area       string    `json:"area" binding:"required"`
	Category   string    `json:"category" binding:"required"`
}

type UpdateJasantar struct {
	// IdJasantar uuid.UUID `json:"-"`
	// IdUser     uuid.UUID `json:"idUser"`
	Title      string    `json:"title"`
	Area       string    `json:"area"`
	Category   string    `json:"category"`
}
