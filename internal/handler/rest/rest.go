package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"braw-api/entity"
	"braw-api/internal/service"
	"braw-api/pkg/middleware"
	"braw-api/pkg/response"
	"github.com/gin-gonic/gin"
	
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint(router *gin.Engine) {
	  
    r.router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }
        c.Next()
    })

    r.router.Use(r.middleware.Timeout())
	
    routerGroup := r.router.Group("/api/v1")
    routerGroup.GET("/health-check", healthCheck)
    routerGroup.GET("/time-out", testTimeout)
    routerGroup.GET("/login-user", r.middleware.AuthenticateUser, getLoginUser)
    routerGroup.POST("/register", r.Register)
    routerGroup.POST("/login", r.Login)

	user := r.router.Group("/user")
	user.POST("/profile/upload", r.middleware.AuthenticateUser, r.UploadPhoto)

	preloved := r.router.Group("/api")
	preloved.POST("/preloved", r.CreatePreloved)
	preloved.GET("/readpreloved/:id", r.GetPrelovedByID)
	preloved.DELETE("/delpreloved/:id", r.middleware.AuthenticateUser, r.middleware.OnlyAdmin, r.DeletePreloved)
	preloved.PATCH("/updatepreloved/:id", r.UpdatePreloved)
	preloved.GET("/", r.GetAllPreloved)
}

func (r *Rest) Serve() {
	//addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("PORT")
	if port == "" {
        port = "5000"
    }

	err := r.router.Run(fmt.Sprintf(":%s",port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	response.Success(ctx, http.StatusOK, "get login user", user.(entity.User))
}