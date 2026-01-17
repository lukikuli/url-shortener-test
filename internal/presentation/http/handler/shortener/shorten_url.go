package shortener

import (
	"doit/urlshortener/internal/presentation/http/request"
	"doit/urlshortener/internal/presentation/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenUrl(c *gin.Context) {
	var payload request.RequestShortenUrl

	if payload.LongUrl == "" {
		response.ErrorResponse(c, http.StatusNotFound, "url is required")
		return
	}

	result, err := h.shortenerUC.ShortenURL(c, payload.LongUrl, payload.TTLSeconds)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, result)
}
