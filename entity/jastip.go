package entity

import (
	"time"

	"github.com/google/uuid"
)

type Jastip struct {
	IdJastip   uuid.UUID `json:"idJastip" gorm:"type:varchar(36);primary_key;"`
	IdUser     uuid.UUID `json:"idUser" gorm:"type:varchar(36);foreign_key;"`
	Title      string    `json:"title" gorm:"type:varchar(255);not null;"`
	Category   string    `json:"category" gorm:"type:varchar(255);not null;"`
	Price      int       `json:"price" gorm:"type:integer;not null;"`
	OpenDay    string    `json:"openDay" gorm:"type:varchar(255);not null;"`
	CloseOrder string `json:"closeOrder" gorm:"type:varchar(10);not null;"`
	//PhotoLink  string    `json:"-" gorm:"type:varchar(200)"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IdPhotolink uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);foreign_key;"`
}
