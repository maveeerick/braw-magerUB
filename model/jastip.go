package model

import (
	"github.com/google/uuid"
)

type CreateJastip struct {
	ID_Jastip   uuid.UUID	`json:"-"`
	ID_User       uuid.UUID `json:"-"`
	Title       string `json:"title" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Price       int    `json:"price,string" binding:"required"`
	Open_Day    string `json:"open_day" binding:"required"`
	Close_Order string`json:"close_order" binding:"required"`
}

type UpdateJastip struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	Price       int    `json:"price,string"`
	Open_Day    string `json:"open_day"`
	Close_Order string `json:"close_order"`
}