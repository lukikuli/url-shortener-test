package response

import "github.com/gin-gonic/gin"

type successResponse struct {
	Status any `json:"status"`
	Data   any `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data any) {
	resp := successResponse{
		Status: map[string]interface{}{
			"code": code,
			"msg":  "success",
		},
		Data: data,
	}
	c.AbortWithStatusJSON(code, resp)
}
