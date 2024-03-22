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
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
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
	user.PATCH("/updatedata/:id", r.UpdateUserData)
	user.GET("/readuserdatabyuserid/:id", r.GetUserDataByUserID)

	preloved := r.router.Group("/api")
	preloved.POST("/preloved", r.CreatePreloved)
	preloved.GET("/readpreloved/:id", r.GetPrelovedByID)
	preloved.DELETE("/delpreloved/:id", r.middleware.AuthenticateUser, r.DeletePreloved)
	preloved.PATCH("/updatepreloved/:id", r.UpdatePreloved)
	preloved.GET("/getallpreloved", r.GetAllPreloved)
	preloved.GET("/readprelovedbyuserid/:id", r.GetPrelovedByUserID)
	preloved.POST("/preloved/:prelovedId/image/upload", r.middleware.AuthenticateUser, r.UploadPrelovedImage)
	preloved.GET("/preloved/:prelovedId/image",  r.GetPrelovedImagesByPrelovedId)
	preloved.GET("/preloved/:prelovedId/image/:imageId", r.GetPrelovedImagesByImageId)
	preloved.DELETE("/preloved/:prelovedId/image/:imageId", r.DeletePrelovedImagesByImageId)

	jastip := r.router.Group("/api")
	jastip.POST("/jastip", r.CreateJastip)
	jastip.GET("/readjastip/:id", r.GetJastipByID)
	jastip.DELETE("/deljastip/:id", r.middleware.AuthenticateUser, r.DeleteJastip)
	jastip.PATCH("/updatejastip/:id", r.UpdateJastip)
	jastip.GET("/getalljastip", r.GetAllJastip)
	jastip.GET("/readjastipbyuserid/:id", r.GetJastipByUserID)
	jastip.POST("/jastip/:jastipId/image/upload", r.middleware.AuthenticateUser, r.UploadJastipImage)
	jastip.GET("/jastip/:jastipId/image", r.GetJastipImagesByJastipId)
	jastip.GET("/jastip/:jastipId/image/:imageId", r.GetJastipImagesByImageId)
	jastip.DELETE("/jastip/:jastipId/image/:imageId", r.DeleteJastipImagesByImageId)

	jasantar := r.router.Group("/api")
	jasantar.POST("/jasantar", r.CreateJasantar)
	jasantar.GET("/readjasantar/:id", r.GetJasantarByID)
	jasantar.DELETE("/deljasantar/:id", r.middleware.AuthenticateUser, r.DeleteJasantar)
	jasantar.PATCH("/updatejasantar/:id", r.UpdateJasantar)
	jasantar.GET("/getalljasantar", r.GetAllJasantar)
	jasantar.GET("/readjasantarbyuserid/:id", r.GetJasantarByUserID)
	jasantar.POST("/jasantar/:jasantarId/image/upload", r.middleware.AuthenticateUser, r.UploadJasantarImage)
	jasantar.GET("/jasantar/:jasantarId/image", r.GetJasantarImagesByJasantarId)
	jasantar.GET("/jasantar/:jasantarId/image/:imageId", r.GetJasantarImagesByImageId)
	jasantar.DELETE("/jasantar/:jasantarId/image/:imageId", r.DeleteJasantarImagesByImageId)

	komunitasbraw := r.router.Group("/api")
	komunitasbraw.POST("/komunitasbraw", r.CreateKomunitasbraw)
	komunitasbraw.GET("/readkomunitasbraw/:id", r.GetKomunitasbrawByID)
	komunitasbraw.DELETE("/delkomunitasbraw/:id", r.middleware.AuthenticateUser, r.DeleteKomunitasbraw)
	komunitasbraw.PATCH("/updatekomunitasbraw/:id", r.UpdateKomunitasbraw)
	komunitasbraw.GET("/getallkomunitasbraw", r.GetAllKomunitasbraw)
	komunitasbraw.GET("/readkomunitasbrawbyuserid/:id", r.GetKomunitasbrawByUserID)
	komunitasbraw.POST("/komunitasbraw/:komunitasbrawId/image/upload", r.middleware.AuthenticateUser, r.UploadKomunitasbrawImage)
	komunitasbraw.GET("/komunitasbraw/:komunitasbrawId/image", r.GetKomunitasbrawImagesByKomunitasbrawId)
	komunitasbraw.GET("/komunitasbraw/:komunitasbrawId/image/:imageId", r.GetKomunitasbrawImagesByImageId)
	komunitasbraw.DELETE("/komunitasbraw/:komunitasbrawId/image/:imageId", r.DeleteKomunitasbrawImagesByImageId)
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