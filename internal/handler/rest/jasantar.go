package rest

import (
	"log"
	"net/http"

	//"strconv"

	"braw-api/model"
	"braw-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateJasantar(ctx *gin.Context) {
	var jasantarReq model.CreateJasantar

	if err := ctx.ShouldBindJSON(&jasantarReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	jasantar, err := r.service.JasantarService.CreateJasantar(&jasantarReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create jasantar", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create jasantar", jasantar)
}

func (r *Rest) GetJasantarByID(ctx *gin.Context) {
	jasantarID := ctx.Param("id")

	jasantar, err := r.service.JasantarService.GetJasantarByID(jasantarID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get jasantar", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get jasantar", jasantar)
}

func (r *Rest) GetJasantarByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")

	jasantar, err := r.service.JasantarService.GetJasantarByUserID(userID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get jasantar", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get jasantar", jasantar)
}

func (r *Rest) DeleteJasantar(ctx *gin.Context) {
	jasantarID := ctx.Param("id")
	log.Println(jasantarID)
	err := r.service.JasantarService.DeleteJasantar(jasantarID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete jasantar", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete jasantar", nil)
}

func (r *Rest) UpdateJasantar(ctx *gin.Context) {
	jasantarID := ctx.Param("id")

	var jasantarReq model.UpdateJasantar
	if err := ctx.ShouldBindJSON(&jasantarReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	jasantar, err := r.service.JasantarService.UpdateJasantar(&jasantarReq, jasantarID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update jasantar", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update jasantar", jasantar)
}

func (r *Rest) GetAllJasantar(ctx *gin.Context) {

	jasantar, err := r.service.JasantarService.GetAllJasantar()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all jasantar", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all jasantar", jasantar)
}

func (r *Rest) UploadJasantarImage(ctx *gin.Context) {
	jasantarID := ctx.Param("jasantarId")
	jasantarUUID, errParse := uuid.Parse(jasantarID)
	if errParse != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jasantar id", errParse)
		return
	}

	// Parse form data to get uploaded files
	file, _ := ctx.FormFile("file")

	image, errServ := r.service.JasantarImagesService.CreateJasantarImage(jasantarUUID, file)
	if errServ != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to upload image", errServ)
		return
	}

	response.Success(ctx, http.StatusOK, "upload image success", image)
}

func (r *Rest) GetJasantarImagesByJasantarId(ctx *gin.Context) {
	jasantarID := ctx.Param("jasantarId")
	if _, err := uuid.Parse(jasantarID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jasantar id", err)
		return
	}

	images, err := r.service.JasantarImagesService.GetImagesByJasantarID(jasantarID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting images", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", images)
}

func (r *Rest) GetJasantarImagesByImageId(ctx *gin.Context) {
	imageId := ctx.Param("imageId")

	jasantarID := ctx.Param("jasantarId")
	if _, err := uuid.Parse(jasantarID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jasantar id", err)
		return
	}

	image, err := r.service.JasantarImagesService.GetImageByID(imageId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", image)
}

func (r *Rest) DeleteJasantarImagesByImageId(ctx *gin.Context) {
	imageId := ctx.Param("imageId")

	jasantarID := ctx.Param("jasantarId")
	if _, err := uuid.Parse(jasantarID); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid jasantar id", err)
		return
	}

	err := r.service.JasantarImagesService.DeleteJasantarImage(imageId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error getting image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete jasantar image", nil)
}