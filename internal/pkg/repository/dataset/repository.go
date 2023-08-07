package dataset

type (
	Repository struct {
		FilePath string
	}
)

func NewRepository(filePath string) *Repository {
	return &Repository{FilePath: filePath}
}
