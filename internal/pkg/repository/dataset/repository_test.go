package dataset

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type datasetSuiteTest struct {
	suite.Suite
	ctx  context.Context
	repo *Repository
}

func (t *datasetSuiteTest) SetupTest() {
	t.ctx = context.Background()
	t.repo = NewRepository("./test_path.csv")
}

func (t *datasetSuiteTest) TestNewRepository() {
	t.NotNil(t.repo)
}

func TestDataset(t *testing.T) {
	suite.Run(t, new(datasetSuiteTest))
}
