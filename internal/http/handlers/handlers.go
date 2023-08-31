package handlers

import "github.com/gin-gonic/gin"

type ArticleHandler interface {
	GetArticles(ctx *gin.Context)
	CreateArticle(ctx *gin.Context)
}
