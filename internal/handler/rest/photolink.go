package rest

import (
	"log"
	"net/http"

	"braw-api/model"
	"braw-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatePhotolink(ctx *gin.Context) {
	var photolinkReq model.CreatePhotolink

	if err := ctx.ShouldBindJSON(&photolinkReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	photolink, err := r.service.PhotolinkService.CreatePhotolink(&photolinkReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create photolink", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create photolink", photolink)
}

func (r *Rest) GetPhotolinkByID(ctx *gin.Context) {
	photolinkID := ctx.Param("id")

	photolink, err := r.service.PhotolinkService.GetPhotolinkByID(photolinkID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get photolink", photolink)
}

func (r *Rest) DeletePhotolink(ctx *gin.Context) {
	photolinkID := ctx.Param("id")
	log.Println(photolinkID)
	err := r.service.PhotolinkService.DeletePhotolink(photolinkID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete photolink", nil)
}

func (r *Rest) UpdatePhotolink(ctx *gin.Context) {
	photolinkID := ctx.Param("id")

	var photolinkReq model.UpdatePhotolink
	if err := ctx.ShouldBindJSON(&photolinkReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	photolink, err := r.service.PhotolinkService.UpdatePhotolink(&photolinkReq, photolinkID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update photolink", photolink)
}

func (r *Rest) GetAllPhotolink(ctx *gin.Context) {
	photolink, err := r.service.PhotolinkService.GetAllPhotolink()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all photolink", photolink)
}
