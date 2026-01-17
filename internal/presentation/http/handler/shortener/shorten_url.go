package shortener

import (
	"doit/urlshortener/internal/presentation/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenUrl(c *gin.Context) {

	result, err := h.shortenerUC.CreateShortUrl(c, "")
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, result)
}
