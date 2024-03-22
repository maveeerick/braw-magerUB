package rest

import (
	"errors"
	"log"
	"net/http"
	"strings"

	//"strconv"

	"braw-api/model"
	"braw-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateJastip(ctx *gin.Context) {
	var jastipReq model.CreateJastip

	if err := ctx.ShouldBindJSON(&jastipReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	jastip, err := r.service.JastipService.CreateJastip(&jastipReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create jastip", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create Jastip", jastip)
}

func (r *Rest) GetJastipByID(ctx *gin.Context) {
	jastipID := ctx.Param("id")

	jastip, err := r.service.JastipService.GetJastipByID(jastipID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get jastip", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get jastip", jastip)
}

func (r *Rest) GetJastipByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")

	jastip, err := r.service.JastipService.GetJastipByUserID(userID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get jastip", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get jastip", jastip)
}

func (r *Rest) DeleteJastip(ctx *gin.Context) {
	jastipID := ctx.Param("id")
	log.Println(jastipID)
	err := r.service.JastipService.DeleteJastip(jastipID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete jastip", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete jastip", nil)
}

func (r *Rest) UpdateJastip(ctx *gin.Context) {
	jastipID := ctx.Param("id")

	var jastipReq model.UpdateJastip
	if err := ctx.ShouldBindJSON(&jastipReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	jastip, err := r.service.JastipService.UpdateJastip(&jastipReq, jastipID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update jastip", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update jastip", jastip)
}

func (r *Rest) GetAllJastip(ctx *gin.Context) {
	jastip, err := r.service.JastipService.GetAllJastip()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all jastip", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all jastip", jastip)
}

func (r *Rest) UploadJastipImage(ctx *gin.Context) {
	jastipID := ctx.Param("jastipId")
	jastipUUID, errParse := uuid.Parse(jastipID)
	if errParse != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jastip id", errParse)
		return
	}

	file, _ := ctx.FormFile("file")

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		response.Error(ctx, http.StatusBadRequest, "failed to upload file", errors.New("only accept Content-Type image/"))
		return
	}

	image, errServ := r.service.JastipImagesService.CreateJastipImage(jastipUUID, file)
	if errServ != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to upload image", errServ)
		return
	}

	response.Success(ctx, http.StatusOK, "upload image success", image)
}

func (r *Rest) GetJastipImagesByJastipId(ctx *gin.Context) {
	jastipID := ctx.Param("jastipId")
	if _, err := uuid.Parse(jastipID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jastip id", err)
		return
	}

	images, err := r.service.JastipImagesService.GetImagesByJastipID(jastipID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting images", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", images)
}

func (r *Rest) GetJastipImagesByImageId(ctx *gin.Context) {
	imageId := ctx.Param("imageId")

	jastipID := ctx.Param("jastipId")
	if _, err := uuid.Parse(jastipID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jastip id", err)
		return
	}

	image, err := r.service.JastipImagesService.GetImageByID(imageId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", image)
}

func (r *Rest) DeleteJastipImagesByImageId(ctx *gin.Context) {
	imageId := ctx.Param("imageId")

	jastipID := ctx.Param("jastipId")
	if _, err := uuid.Parse(jastipID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jastip id", err)
		return
	}

	err := r.service.JastipImagesService.DeleteJastipImage(imageId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete jastip image", nil)
}