package rest

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"braw-api/model"
	"braw-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateKomunitasbraw(ctx *gin.Context) {
	var komunitasbrawReq model.CreateKomunitasbraw

	if err := ctx.ShouldBindJSON(&komunitasbrawReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	komunitasbraw, err := r.service.KomunitasbrawService.CreateKomunitasbraw(&komunitasbrawReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create komunitasbraw", komunitasbraw)
}

func (r *Rest) GetKomunitasbrawByID(ctx *gin.Context) {
	komunitasbrawID := ctx.Param("id")

	komunitasbraw, err := r.service.KomunitasbrawService.GetKomunitasbrawByID(komunitasbrawID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get komunitasbraw", komunitasbraw)
}

func (r *Rest) GetKomunitasbrawByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")

	komunitasbraw, err := r.service.KomunitasbrawService.GetKomunitasbrawByUserID(userID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get komunitasbraw", komunitasbraw)
}

func (r *Rest) DeleteKomunitasbraw(ctx *gin.Context) {
	komunitasbrawID := ctx.Param("id")
	log.Println(komunitasbrawID)
	err := r.service.KomunitasbrawService.DeleteKomunitasbraw(komunitasbrawID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete komunitasbraw", nil)
}

func (r *Rest) UpdateKomunitasbraw(ctx *gin.Context) {
	komunitasbrawID := ctx.Param("id")

	var komunitasbrawReq model.UpdateKomunitasbraw
	if err := ctx.ShouldBindJSON(&komunitasbrawReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	komunitasbraw, err := r.service.KomunitasbrawService.UpdateKomunitasbraw(&komunitasbrawReq, komunitasbrawID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update komunitasbraw", komunitasbraw)
}

func (r *Rest) GetAllKomunitasbraw(ctx *gin.Context) {
	komunitasbraw, err := r.service.KomunitasbrawService.GetAllKomunitasbraw()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all komunitasbraw", komunitasbraw)
}

func (r *Rest) UploadKomunitasbrawImage(ctx *gin.Context) {
	komunitasbrawID := ctx.Param("komunitasbrawId")
	komunitasbrawUUID, errParse := uuid.Parse(komunitasbrawID)
	if errParse != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid komunitasbraw id", errParse)
		return
	}

	file, _ := ctx.FormFile("file")

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		response.Error(ctx, http.StatusBadRequest, "failed to upload file", errors.New("only accept Content-Type image/"))
		return
	}

	image, errServ := r.service.KomunitasbrawImagesService.CreateKomunitasbrawImage(komunitasbrawUUID, file)
	if errServ != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to upload image", errServ)
		return
	}

	response.Success(ctx, http.StatusOK, "upload image success", image)
}

func (r *Rest) GetKomunitasbrawImagesByKomunitasbrawId(ctx *gin.Context) {
	komunitasbrawID := ctx.Param("komunitasbrawId")
	if _, err := uuid.Parse(komunitasbrawID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid komunitasbraw id", err)
		return
	}

	images, err := r.service.KomunitasbrawImagesService.GetImagesByKomunitasbrawID(komunitasbrawID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting images", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", images)
}

func (r *Rest) GetKomunitasbrawImagesByImageId(ctx *gin.Context) {
	imageId := ctx.Param("imageId")

	komunitasbrawID := ctx.Param("komunitasbrawId")
	if _, err := uuid.Parse(komunitasbrawID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid komunitasbraw id", err)
		return
	}

	image, err := r.service.KomunitasbrawImagesService.GetImageByID(imageId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", image)
}

func (r *Rest) DeleteKomunitasbrawImagesByImageId(ctx *gin.Context) {
	imageId := ctx.Param("imageId")

	komunitasbrawID := ctx.Param("komunitasbrawId")
	if _, err := uuid.Parse(komunitasbrawID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid komunitasbraw id", err)
		return
	}

	err := r.service.KomunitasbrawImagesService.DeleteKomunitasbrawImage(imageId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete komunitasbraw image", nil)
}