package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegister struct {
	IdUser   uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	Username string    `json:"username" binding:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserParam struct {
	IdUser   uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type UserUploadPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}
