package rest

import (
	"braw-api/model"
	"braw-api/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatePreloved(ctx *gin.Context) {
	var prelovedReq model.CreatePreloved

	if err := ctx.ShouldBindJSON(&prelovedReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	preloved, err := r.service.PrelovedService.CreatePreloved(&prelovedReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create preloved", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create preloved", preloved)
}

func (r *Rest) GetPrelovedByID(ctx *gin.Context) {
	prelovedID := ctx.Param("id")

	preloved, err := r.service.PrelovedService.GetPrelovedByID(prelovedID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get preloved", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get preloved", preloved)
}

func (r *Rest) GetPrelovedByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")

	preloved, err := r.service.PrelovedService.GetPrelovedByUserID(userID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get preloved", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get preloved", preloved)
}

func (r *Rest) DeletePreloved(ctx *gin.Context) {
	prelovedID := ctx.Param("id")
	log.Println(prelovedID)
	err := r.service.PrelovedService.DeletePreloved(prelovedID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete preloved", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete preloved", nil)
}

func (r *Rest) UpdatePreloved(ctx *gin.Context) {
	prelovedID := ctx.Param("id")

	var prelovedReq model.UpdatePreloved
	if err := ctx.ShouldBindJSON(&prelovedReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	preloved, err := r.service.PrelovedService.UpdatePreloved(&prelovedReq, prelovedID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update preloved", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update preloved", preloved)
}

func (r *Rest) GetAllPreloved(ctx *gin.Context) {
	preloved, err := r.service.PrelovedService.GetAllPreloved()

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all preloved", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all preloved", preloved)
}
