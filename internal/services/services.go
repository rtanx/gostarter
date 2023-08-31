package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rtanx/gostarter/internal/http/dto/requests"
	"github.com/rtanx/gostarter/internal/model"
)

type ArticleService interface {
	Create(ctx *gin.Context, data *requests.ArticleForm) (model.Article, error)
	Get(ctx *gin.Context) (model.Article, error)
}
