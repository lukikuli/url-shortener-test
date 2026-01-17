package shortener

import (
	"doit/urlshortener/internal/presentation/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStats(c *gin.Context) {
	shortCode := c.Param("short_code")

	stats, err := h.shortenerUC.GetStats(c.Request.Context(), shortCode)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, stats)
}
