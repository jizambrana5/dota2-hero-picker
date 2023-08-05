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
	}
	Storage interface {
		GetHero(ctx context.Context, id string) (domain.Hero, error)
		GetAllHeroes(ctx context.Context) ([]domain.Hero, error)
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

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}
