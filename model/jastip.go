package model

import (
	"github.com/google/uuid"
)

type CreateJastip struct {
	IdJastip   uuid.UUID `json:"-"`
	IdUser     uuid.UUID `json:"idUser" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Category   string    `json:"category" binding:"required"`
	Price      int       `json:"price,string" binding:"required"`
	OpenDay    string    `json:"openDay" binding:"required"`
	CloseOrder string    `json:"closeOrder" binding:"required"`
}

type UpdateJastip struct {
	Title      string `json:"title"`
	Category   string `json:"category"`
	Price      int    `json:"price,string"`
	OpenDay    string `json:"openDay"`
	CloseOrder string `json:"closeOrder"`
}
