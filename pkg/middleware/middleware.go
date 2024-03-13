package middleware

import (
	"braw-api/internal/service"
	"braw-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	Timeout() gin.HandlerFunc
	OnlyAdmin(ctx *gin.Context)
	CORSMiddleware() gin.HandlerFunc
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

func (m *middleware) CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Adjust origin as needed
        c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers, Cache-Control, Content-Type")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}