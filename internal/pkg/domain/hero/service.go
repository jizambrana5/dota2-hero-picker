package hero

import (
	"context"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

type (
	Service struct {
		storage Storage
		dataset Dataset
	}
	Storage interface {
		GetHero(ctx context.Context, id string) (domain.Hero, error)
		GetAllHeroes(ctx context.Context) ([]domain.Hero, error)
		SaveHero(ctx context.Context, hero domain.Hero) error
	}
	Dataset interface {
		GetRecords(ctx context.Context) ([][]string, error)
	}
)

func NewService(storage Storage, dataset Dataset) *Service {
	return &Service{storage: storage, dataset: dataset}
}
