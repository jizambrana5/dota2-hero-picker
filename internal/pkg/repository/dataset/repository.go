package dataset

import (
	"log"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
)

type (
	Repository struct {
		FilePath string
	}
)

var _ hero.Dataset = (*Repository)(nil)

func NewRepository(filePath string) *Repository {
	if filePath == "" {
		log.Panic("Empty dataset file path")
	}
	return &Repository{FilePath: filePath}
}
