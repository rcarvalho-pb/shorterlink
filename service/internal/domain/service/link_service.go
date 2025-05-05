package service

import (
	"github.com/rcarvalho-pb/shortlinker-go/internal/domain/entity"
	"github.com/rcarvalho-pb/shortlinker-go/internal/domain/port"
)

type LinkService struct {
	repo port.Repository
}

func (ls *LinkService) CreateLink(req *entity.CreateLinkRequest) error {
	return nil
}

func createSlug(url string) string {
	return ""
}
