package repositories

import (
	"context"

	"github.com/rtanx/gostarter/internal/model"
)

type ArticleRepository interface {
	Create(ctx context.Context, src *model.Article) error
	Get(ctx context.Context, dst []*model.Article) error
}
