package dataset

import (
	"context"
	"encoding/csv"
	"os"
)

func (r Repository) GetRecords(ctx context.Context) ([][]string, error) {
	fd, err := os.Open(r.FilePath)
	if err != nil {
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
