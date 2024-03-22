package model

import "github.com/google/uuid"

type CreatePrelovedPhotos struct {
	IdPrelovedPhoto uuid.UUID `json:"-"`
	PhotoLink       string    `json:"photoLink" binding:"required"`
}

type CreateJastipPhotos struct {
	IdJastipPhoto uuid.UUID `json:"-"`
	PhotoLink     string    `json:"photoLink" binding:"required"`
}

type CreateJasantarPhotos struct {
	IdJasantarPhotolink uuid.UUID `json:"-"`
	PhotoLink           string    `json:"photoLink" binding:"required"`
}

type CreateKomunitasbrawPhotos struct {
	IdKomunitasbrawPhotolink uuid.UUID `json:"-"`
	PhotoLink                string    `json:"photoLink" binding:"required"`
}
