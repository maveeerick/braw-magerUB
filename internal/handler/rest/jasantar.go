package rest

import (
	"log"
	"net/http"
	//"strconv"

	"braw-api/model"
	"braw-api/pkg/response"
	"github.com/gin-gonic/gin"
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