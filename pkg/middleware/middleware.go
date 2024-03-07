package middleware

import (
	"braw-api/internal/service"
	"braw-api/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	Timeout() gin.HandlerFunc
	OnlyAdmin(ctx *gin.Context)
}

type middleware struct {
	jwtAuth jwt.Interface
	service *service.Service
}

func Init(jwtAuth jwt.Interface, service *service.Service) *middleware {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}