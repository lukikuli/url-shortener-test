package route

import (
	shortenerUC "doit/urlshortener/internal/application/usecase/shortener"
	"doit/urlshortener/internal/presentation/http/handler/shortener"

	"github.com/gin-gonic/gin"
)

func InitShortenRoute(route *gin.Engine, uc shortenerUC.ShortenerUsecase) {
	var shortenerHandler = shortener.NewShortenerHandler(uc)

	route.POST("/s/shorten", shortenerHandler.ShortenUrl)
	route.GET("/s/:short_code", shortenerHandler.RedirectUrl)
}
