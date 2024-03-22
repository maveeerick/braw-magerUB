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

func (r *Rest) UpdateUserData(ctx *gin.Context) {
	userID := ctx.Param("id")

	var userDataReq model.UpdateUserData
	if err := ctx.ShouldBindJSON(&userDataReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	userData, err := r.service.UserService.UpdateUserData(&userDataReq, userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update user", userData)
}

func (r *Rest) GetUserDataByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")

	userData, err := r.service.UserService.GetUserDataByUserID(userID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get user data", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get user data", userData)
}