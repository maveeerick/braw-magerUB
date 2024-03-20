package entity

import "github.com/google/uuid"

type Photolink struct {
	IdPhotolink uuid.UUID `json:"idPhotolink" gorm:"type:varchar(36);primary_key;"`
	PhotoLink   string    `json:"-" gorm:"type:varchar(200)"`
	//IdUser      uuid.UUID `json:"idUser" gorm:"type:varchar(36);"`
	// IdPreloved  uuid.UUID `json:"idPreloved" gorm:"type:varchar(36);foreign_key;"`
	// IdJastip    uuid.UUID `json:"idJastip" gorm:"type:varchar(36);foreign_key;"`
	// IdJasantar  uuid.UUID `json:"idJasantar" gorm:"type:varchar(36);foreign_key;"`
	// IdKomunitasbraw uuid.UUID `json:"idKomunitasbraw" gorm:"type:varchar(36);foreign_key;"`
}
