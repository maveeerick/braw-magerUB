package entity

import (
	"time"

	"github.com/google/uuid"
)

type Jasantar struct {
	ID_Jasantar uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	ID_User     uuid.UUID    `json:"id_user" gorm:"type:varchar(36);foreign_key;"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null;"`
	Area        string    `json:"area" gorm:"type:varchar(255);not null;"`
	Category    string    `json:"category" gorm:"type:varchar(255);not null;"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	
}