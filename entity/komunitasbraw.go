package entity

import "github.com/google/uuid"

type Komunitasbraw struct {
	IdKomunitasbraw uuid.UUID `json:"idKomunitasbraw" gorm:"type:varchar(36);primary_key;"`
	IdUser          uuid.UUID `json:"idUser" gorm:"type:varchar(36);foreign_key;"`
	Title           string    `json:"title" gorm:"type:varchar(255);not null;"`
	Category        string    `json:"category" gorm:"type:varchar(255);not null;"`
	Description     string    `json:"description" gorm:"type:varchar(255)"`
	IdPhotolink     uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);foreign_key;"`
}
