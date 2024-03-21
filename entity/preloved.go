package entity

import (
	"time"

	"github.com/google/uuid"
)

type Preloved struct {
	IdPreloved  uuid.UUID `json:"idPreloved" gorm:"type:varchar(36);primary_key;"`
	IdUser      uuid.UUID `json:"idUser" gorm:"type:varchar(36);not null;foreign_key;"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null;"`
	Category    string    `json:"category" gorm:"type:varchar(255);not null;"`
	Price       int       `json:"price" gorm:"type:integer;not null;"`
	Condition   string    `json:"condition" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	IdPhotolink uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);foreign_key;"`
	//PhotoLink   string    `json:"-" gorm:"type:varchar(200)"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
