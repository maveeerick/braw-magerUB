package entity

import "github.com/google/uuid"

type PrelovedPhotos struct {
	IdPrelovedPhoto uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);primary_key;"`
	IdPreloved uuid.UUID `json:"idPreloved" gorm:"type:varchar(36);not null;foreign_key;"`
	PhotoLink       string    `json:"-" gorm:"type:varchar(200)"`
}
