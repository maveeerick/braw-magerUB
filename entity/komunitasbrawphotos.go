package entity

import "github.com/google/uuid"

type KomunitasbrawPhotos struct {
	IdKomunitasbrawphoto uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);primary_key;"`
	IdKomunitasbraw      uuid.UUID `json:"idKomunitasbraw" gorm:"type:varchar(36);not null;foreign_key;"`
	PhotoLink            string    `json:"-" gorm:"type:varchar(200)"`
}
