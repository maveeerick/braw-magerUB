package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null;unique"`      
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;"`
	Role      int       `json:"role" gorm:"foreinKey:ID; references:roles; not null;"`
	PhotoLink string    `json:"-" gorm:"type:varchar(200)"`
	Alamat	  string	`json:"alamat" gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}