package repositories

import (
	"context"

	"github.com/rtanx/gostarter/internal/db"
	"github.com/rtanx/gostarter/internal/model"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	dbc := db.GetConn()
	return &articleRepository{
		db: dbc,
	}
}

// Create implements ArticleRepository.
func (r *articleRepository) Create(ctx context.Context, src *model.Article) error {
	panic("unimplemented")
}

// Get implements ArticleRepository.
func (r *articleRepository) Get(ctx context.Context, dst []*model.Article) error {
	panic("unimplemented")
}
