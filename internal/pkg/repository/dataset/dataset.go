package dataset

import (
	"context"
	"encoding/csv"
	"os"

	"go.uber.org/zap"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/logs"
)

func (r Repository) GetRecords(ctx context.Context) ([][]string, error) {
	fd, err := os.Open(r.FilePath)
	if err != nil {
		logs.Logger.Info("error", zap.Error(err))
		return nil, err
	}
	defer fd.Close()
	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
