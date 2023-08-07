package dataset

func (t *datasetSuiteTest) TestGetRecords_OpenError() {
	records, err := t.repo.GetRecords(t.ctx)
	t.NotNil(err)
	t.Nil(records)
}

func (t *datasetSuiteTest) TestGetRecords_Success() {
	t.repo.FilePath = "./dataset.csv"
	records, err := t.repo.GetRecords(t.ctx)
	t.Nil(err)
	t.NotNil(records)
}
