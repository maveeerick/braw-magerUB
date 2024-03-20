package rest

import (
	"log"
	"net/http"

	"braw-api/model"
	"braw-api/pkg/response"

	"github.com/gin-gonic/gin"
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
	// pageQuery := ctx.Query("page")
	// page, err := strconv.Atoi(pageQuery)
	// if err != nil {
	// 	response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
	// }

	komunitasbraw, err := r.service.KomunitasbrawService.GetAllKomunitasbraw()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all komunitasbraw", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all komunitasbraw", komunitasbraw)
}
