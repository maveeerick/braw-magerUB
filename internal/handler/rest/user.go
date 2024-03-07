package rest

import (
	"net/http"

	"braw-api/model"
	"braw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.Register(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success login to system", token)
}


func (r *Rest) UploadPhoto(ctx *gin.Context) {
	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.UploadPhoto(ctx, model.UserUploadPhoto{Photo: photo})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success upload photo", nil)
}