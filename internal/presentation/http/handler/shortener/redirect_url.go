package shortener

import (
	"doit/urlshortener/internal/presentation/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RedirectUrl(c *gin.Context) {

	shortCode := c.Param("short_code")
	if shortCode == "" {
		response.ErrorResponse(c, http.StatusNotFound, "url not found")
		return
	}

	result, err := h.shortenerUC.GetOriginalUrl(c, shortCode)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusFound, result)
}
