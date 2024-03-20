package model

import "github.com/google/uuid"

type CreatePhotolink struct {
	IdPhotolink uuid.UUID `json:"-"`
	PhotoLink   string    `json:"photoLink" binding:"required"`
	//IdUser      uuid.UUID `json:"idUser" gorm:"type:varchar(36);"`
	// IdPreloved uuid.UUID `json:"-"`
	// IdJastip   uuid.UUID `json:"-"`
	// IdJasantar uuid.UUID `json:"-"`
	// IdKomunitasbraw uuid.UUID `json:"-"`
}

type UpdatePhotolink struct {
	IdPhotolink uuid.UUID `json:"-"`
	PhotoLink   string    `json:"photoLink" binding:"required"`
	//IdUser      uuid.UUID `json:"idUser" gorm:"type:varchar(36);"`
	// IdPreloved uuid.UUID `json:"-"`
	// IdJastip   uuid.UUID `json:"-"`
	// IdJasantar uuid.UUID `json:"-"`
	// IdKomunitasbraw uuid.UUID `json:"-"`
}