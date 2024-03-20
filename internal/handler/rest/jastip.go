package rest

import (
	"log"
	"net/http"
	//"strconv"

	"braw-api/model"
	"braw-api/pkg/response"
	"github.com/gin-gonic/gin"
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
	//itemQuery := ctx.Query("item")
	//item, err := strconv.Atoi(itemQuery)
	// if err != nil {
	// 	response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
	// }

	jastip, err := r.service.JastipService.GetAllJastip()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all jastip", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all jastip", jastip)
}