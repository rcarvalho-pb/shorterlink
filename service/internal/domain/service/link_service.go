package service

import (
	"github.com/rcarvalho-pb/shortlinker-go/internal/domain/port"
)

type LinkService struct {
	repo port.Repository
}

func (ls *LinkService) CreateShortLink(url string) string {
	return ""
}
