package hero

import (
	"context"
	"errors"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

func (t *heroSuite) Test_GetHero_StorageError() {
	t.storageMock.GetHeroFunc = func(ctx context.Context, id string) (domain.Hero, error) {
		return domain.Hero{}, errors.New("error getting geto from db")
	}

	h, err := t.service.GetHero(t.ctx, "id_test")
	t.NotNil(err)
	t.Empty(h)
}

func (t *heroSuite) Test_GetHero_Success() {
	h, err := t.service.GetHero(t.ctx, "id_test")
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *heroSuite) Test_GetAllHeroes() {
	h, err := t.service.GetAllHeroes(t.ctx)
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *heroSuite) Test_GetDataSet() {
	h, err := t.service.GetDataSet(t.ctx)
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *heroSuite) Test_SaveHeroes_GetRecordsError() {
	t.datasetMock.GetRecordsFunc = func(ctx context.Context) ([][]string, error) {
		return [][]string{}, errors.New("error getting the dataset")
	}
	err := t.service.SaveHeroes(t.ctx)
	t.NotNil(err)
}

func (t *heroSuite) Test_SaveHeroes_SaveHeroError() {
	t.storageMock.SaveHeroFunc = func(ctx context.Context, hero domain.Hero) error {
		return errors.New("error savin hero in db")
	}
	err := t.service.SaveHeroes(t.ctx)
	t.NotNil(err)
}

func (t *heroSuite) Test_SaveHeroes() {
	err := t.service.SaveHeroes(t.ctx)
	t.Nil(err)
}

func (t *heroSuite) Test_GetHeroSuggestion_GetAllHeroesError() {
	t.storageMock.GetAllHeroesFunc = func(ctx context.Context) ([]domain.Hero, error) {
		return []domain.Hero{}, errors.New("error getting all heroes")
	}

	h, err := t.service.GetHeroSuggestion(t.ctx, domain.UserPreferences{
		PrimaryAttribute: "str",
		Roles:            []domain.Role{domain.Support},
	})

	t.Empty(h)
	t.NotNil(err)
}

func (t *heroSuite) Test_GetHeroSuggestion() {

	h, err := t.service.GetHeroSuggestion(t.ctx, domain.UserPreferences{
		PrimaryAttribute: "str",
		Roles:            []domain.Role{domain.Support},
	})

	t.NotEmpty(h)
	t.Nil(err)
}

func (t *heroSuite) Test_FilterHeroes_HasNoRoles() {
	up := domain.UserPreferences{
		PrimaryAttribute: "str",
		Roles:            []domain.Role{domain.Support},
	}

	heroes := []domain.Hero{{
		HeroIndex:        "2",
		PrimaryAttribute: "str",
		NameInGame:       "Ursa",
		Role:             []domain.Role{domain.Carry},
	}}

	filteredHeroes := filterHeroes(up, heroes)
	t.Len(filteredHeroes, 0)
}
