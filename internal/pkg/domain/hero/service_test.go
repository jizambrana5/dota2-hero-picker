package hero

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero/mocks"
)

type heroSuite struct {
	suite.Suite
	ctx         context.Context
	storageMock *mocks.StorageMock
	datasetMock *mocks.DatasetMock
	service     *Service
}

func (t *heroSuite) SetupTest() {
	t.ctx = context.Background()
	t.storageMock = &mocks.StorageMock{
		GetAllHeroesFunc: func(ctx context.Context) ([]domain.Hero, error) {
			return []domain.Hero{{
				HeroIndex:        "1",
				PrimaryAttribute: "str",
				NameInGame:       "Abbadon",
				Role:             []domain.Role{"Support"},
			}}, nil
		},
		GetHeroFunc: func(ctx context.Context, id string) (domain.Hero, error) {
			return domain.Hero{
				HeroIndex:        "id_test",
				PrimaryAttribute: "str	",
				NameInGame:       "Abbadon",
				Role:             []domain.Role{"Support"},
				WinRates:         []domain.Rank{{Name: domain.Herald, Rate: 18.00}},
			}, nil
		},
		SaveHeroFunc: func(ctx context.Context, hero domain.Hero) error {
			return nil
		},
	}
	t.datasetMock = &mocks.DatasetMock{
		GetRecordsFunc: func(ctx context.Context) ([][]string, error) {
			return [][]string{
				{"ID", "Name", "Primary Attribute", "Roles", "", "", "", "", "", "", "", "", "", ""},
				{"1", "Abbado", "str", "Support", "", "", "45.00", "", "", "45.10", "", "", "45.20", "", "", "45.30", "", "", "45.40", "", "", "45.50"},
			}, nil
		}}
	t.service = NewService(t.storageMock, t.datasetMock)
}

func (t *heroSuite) Test_NewService() {
	t.NotNil(NewService(t.storageMock, t.datasetMock))
}

func TestHero(t *testing.T) {
	suite.Run(t, new(heroSuite))
}
