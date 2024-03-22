package entity

import "github.com/google/uuid"

type JasantarPhotos struct {
	IdJasantarPhoto uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);primary_key;"`
	PhotoLink   string    `json:"-" gorm:"type:varchar(200)"`
}
