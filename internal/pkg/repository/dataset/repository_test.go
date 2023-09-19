package dataset

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/logs"
)

type datasetSuiteTest struct {
	suite.Suite
	ctx  context.Context
	repo *Repository
}

func (t *datasetSuiteTest) SetupTest() {
	// Initialize the logger based on the environment
	logs.InitLogger("development")
	t.ctx = context.Background()
	t.repo = NewRepository("./test_path.csv")
}

func (t *datasetSuiteTest) TestNewRepository() {
	t.NotNil(t.repo)
}

func TestDataset(t *testing.T) {
	suite.Run(t, new(datasetSuiteTest))
}
