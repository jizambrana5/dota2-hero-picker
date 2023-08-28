//go:generate moq -pkg mocks -out ./mocks/hero_mocks.go -skip-ensure . Storage Dataset
package hero

import (
	"context"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/rest"
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

var _ rest.HeroService = (*Service)(nil)

func NewService(storage Storage, dataset Dataset) *Service {
	return &Service{storage: storage, dataset: dataset}
}
