package rest

import (
	"log"
	"net/http"

	"braw-api/model"
	"braw-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatePrelovedPhotos(ctx *gin.Context) {
	var prelovedPhotosReq model.CreatePrelovedPhotos

	if err := ctx.ShouldBindJSON(&prelovedPhotosReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	prelovedPhotos, err := r.service.PhotolinkService.CreatePrelovedPhotos(&prelovedPhotosReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create photolink", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create photolink", prelovedPhotos)
}

func (r *Rest) GetPrelovedPhotosByID(ctx *gin.Context) {
	prelovedPhotosID := ctx.Param("id")

	prelovedPhotos, err := r.service.PhotolinkService.GetPrelovedPhotosByID(prelovedPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get photolink", prelovedPhotos)
}

func (r *Rest) DeletePrelovedPhotos(ctx *gin.Context) {
	prelovedPhotosID := ctx.Param("id")
	log.Println(prelovedPhotosID)
	err := r.service.PhotolinkService.DeletePrelovedPhotos(prelovedPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete photolink", nil)
}

func (r *Rest) CreateJastipPhotos(ctx *gin.Context) {
	var jastipPhotosReq model.CreateJastipPhotos

	if err := ctx.ShouldBindJSON(&jastipPhotosReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	jastipPhotos, err := r.service.PhotolinkService.CreateJastipPhotos(&jastipPhotosReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create photolink", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create photolink", jastipPhotos)
}

func (r *Rest) GetJastipPhotosByID(ctx *gin.Context) {
	jastipPhotosID := ctx.Param("id")

	jastipPhotos, err := r.service.PhotolinkService.GetJastipPhotosByID(jastipPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get photolink", jastipPhotos)
}

func (r *Rest) DeleteJastipPhotos(ctx *gin.Context) {
	jastipPhotosID := ctx.Param("id")
	log.Println(jastipPhotosID)
	err := r.service.PhotolinkService.DeleteJastipPhotos(jastipPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete photolink", nil)
}

func (r *Rest) CreateJasantarPhotos(ctx *gin.Context) {
	var jasantarPhotosReq model.CreateJasantarPhotos

	if err := ctx.ShouldBindJSON(&jasantarPhotosReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	jasantarPhotos, err := r.service.PhotolinkService.CreateJasantarPhotos(&jasantarPhotosReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create photolink", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create photolink", jasantarPhotos)
}

func (r *Rest) GetJasantarPhotosByID(ctx *gin.Context) {
	jasantarPhotosID := ctx.Param("id")

	jasantarPhotos, err := r.service.PhotolinkService.GetJasantarPhotosByID(jasantarPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get photolink", jasantarPhotos)
}

func (r *Rest) DeleteJasantarPhotos(ctx *gin.Context) {
	jasantarPhotosID := ctx.Param("id")
	log.Println(jasantarPhotosID)
	err := r.service.PhotolinkService.DeleteJasantarPhotos(jasantarPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete photolink", nil)
}

func (r *Rest) CreateKomunitasbrawPhotos(ctx *gin.Context) {
	var komunitasbrawPhotosReq model.CreateKomunitasbrawPhotos

	if err := ctx.ShouldBindJSON(&komunitasbrawPhotosReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	komunitasbrawPhotos, err := r.service.PhotolinkService.CreateKomunitasbrawPhotos(&komunitasbrawPhotosReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create photolink", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create photolink", komunitasbrawPhotos)
}

func (r *Rest) GetKomunitasbrawPhotosByID(ctx *gin.Context) {
	komunitasbrawPhotosID := ctx.Param("id")

	komunitasbrawPhotos, err := r.service.PhotolinkService.GetKomunitasbrawPhotosByID(komunitasbrawPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get photolink", komunitasbrawPhotos)
}

func (r *Rest) DeleteKomunitasbrawPhotos(ctx *gin.Context) {
	komunitasbrawPhotosID := ctx.Param("id")
	log.Println(komunitasbrawPhotosID)
	err := r.service.PhotolinkService.DeleteKomunitasbrawPhotos(komunitasbrawPhotosID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete photolink", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete photolink", nil)
}