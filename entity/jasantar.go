package entity

import (
	"time"

	"github.com/google/uuid"
)

type Jasantar struct {
	IdJasantar uuid.UUID `json:"idJasantar" gorm:"type:varchar(36);primary_key;"`
	IdUser     uuid.UUID `json:"idUser" gorm:"type:varchar(36);foreign_key;"`
	Title      string    `json:"title" gorm:"type:varchar(255);not null;"`
	Area       string    `json:"area" gorm:"type:varchar(255);not null;"`
	Category   string    `json:"category" gorm:"type:varchar(255);not null;"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	IdPhotolink uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);foreign_key;"`
}
