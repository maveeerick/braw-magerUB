package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
	


type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"is_success"`
}

func Success(ctx *gin.Context, code int, message string, data any) {
	resp := Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, resp)
}

func Error(ctx *gin.Context, code int, message string, err error) {
	resp := Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	}
	ctx.JSON(code, resp)
}

// Add a middleware function to handle CORS
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Adjust origin as needed
        c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

// Usage in your Gin router
func main() {
    router := gin.Default()

    // Apply CORS middleware globally
    router.Use(CORSMiddleware())

    // Your API routes here
    router.GET("/api/example", func(c *gin.Context) {
        Success(c, http.StatusOK, "Success!", "This is an example response")
    })

    router.Run(":8080") // Adjust port as needed
}