package response

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error" example:"message"`
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errorResponse{msg})
}
