package port

import "github.com/rcarvalho-pb/shortlinker-go/internal/domain/entity"

type Repository interface {
	Add(*entity.Link) error
	Get(string) (*entity.Link, error)
}
