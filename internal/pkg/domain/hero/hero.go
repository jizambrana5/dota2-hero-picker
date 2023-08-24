package hero

import (
	"context"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

type Hero interface {
	GetHero(ctx context.Context, id string) (domain.Hero, error)
	GetAllHeroes(ctx context.Context) ([]domain.Hero, error)
	GetHeroSuggestion(ctx context.Context, preferences domain.UserPreferences) (domain.Hero, error)
}

type (
	Service struct {
		storage Storage
		dataset Dataset
	}
	Storage interface {
		GetHero(ctx context.Context, id string) (domain.Hero, error)
		GetAllHeroes(ctx context.Context) ([]domain.Hero, error)
	}
	Dataset interface {
		GetRecords(ctx context.Context) ([][]string, error)
	}
)

func (s Service) GetAllHeroes(ctx context.Context) ([]Hero, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetHeroSuggestion(ctx context.Context) ([]Hero, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetDataSet(ctx context.Context) ([][]string, error) {
	return s.dataset.GetRecords(ctx)
}

func NewService(storage Storage, dataset Dataset) *Service {
	return &Service{storage: storage, dataset: dataset}
}
