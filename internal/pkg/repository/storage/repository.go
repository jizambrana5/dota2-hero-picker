package storage

import (
	"context"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

type Repository struct{}

func (Repository) GetHero(ctx context.Context, id string) (domain.Hero, error) {
	//TODO implement me
	panic("implement me")
}

func (Repository) GetAllHeroes(ctx context.Context) ([]domain.Hero, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository() Repository {
	return Repository{}
}
