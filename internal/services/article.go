package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rtanx/gostarter/internal/http/dto/requests"
	"github.com/rtanx/gostarter/internal/model"
	"github.com/rtanx/gostarter/internal/repositories"
)

type articleService struct {
	articleRepo repositories.ArticleRepository
}

// Create implements ArticleService.
func (r *articleService) Create(ctx *gin.Context, data *requests.ArticleForm) (model.Article, error) {
	panic("unimplemented")
}

// Get implements ArticleService.
func (r *articleService) Get(ctx *gin.Context) (model.Article, error) {
	panic("unimplemented")
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return &articleService{
		articleRepo,
	}
}
