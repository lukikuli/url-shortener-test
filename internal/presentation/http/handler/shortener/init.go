package shortener

import "doit/urlshortener/internal/application/usecase/shortener"

type Handler struct {
	shortenerUC shortener.ShortenerUsecase
}

func NewShortenerHandler(ucshortenerUC shortener.ShortenerUsecase) *Handler {
	return &Handler{
		shortenerUC: ucshortenerUC,
	}
}
