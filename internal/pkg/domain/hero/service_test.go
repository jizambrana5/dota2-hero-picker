package hero

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero/mocks"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/logs"
)

type heroSuite struct {
	suite.Suite
	ctx         context.Context
	storageMock *mocks.StorageMock
	datasetMock *mocks.DatasetMock
	benchmark   *mocks.BenchmarkMock
	service     *Service
}

func (t *heroSuite) SetupTest() {
	logs.InitLogger("development")
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
				{"1", "Abbadon", "str", "Support", "", "", "45.00", "", "", "45.10", "", "", "45.20", "", "", "45.30", "", "", "45.40", "", "", "45.50"},
				{"2", "Ursa", "agi", "Carry", "", "", "45.00", "", "", "45.10", "", "", "45.20", "", "", "45.30", "", "", "45.40", "", "", "45.50"},
				{"3", "Shaker", "str", "Support", "", "", "45.00", "", "", "45.10", "", "", "45.20", "", "", "45.30", "", "", "45.40", "", "", "45.50"},
			}, nil
		},
	}
	t.benchmark = &mocks.BenchmarkMock{}
	t.service = NewService(t.storageMock, t.datasetMock, t.benchmark)
}

func (t *heroSuite) Test_NewService() {
	t.NotNil(NewService(t.storageMock, t.datasetMock, t.benchmark))
}

func TestHero(t *testing.T) {
	suite.Run(t, new(heroSuite))
}
