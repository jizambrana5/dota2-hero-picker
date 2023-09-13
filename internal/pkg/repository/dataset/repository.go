package dataset

import (
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
)

type (
	Repository struct {
		FilePath string
	}
)

var _ hero.Dataset = (*Repository)(nil)

func NewRepository(filePath string) *Repository {
	return &Repository{FilePath: filePath}
}
