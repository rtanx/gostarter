package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rtanx/gostarter/internal/errs"
	"github.com/rtanx/gostarter/internal/http/dto/requests"
	"github.com/rtanx/gostarter/internal/http/dto/responses"
	"github.com/rtanx/gostarter/internal/infrastructure/logger"
	"github.com/rtanx/gostarter/internal/services"
)

type articleHandler struct {
	articleSrv services.ArticleService
}

func NewArticleHandler(articleSrv services.ArticleService) ArticleHandler {
	return &articleHandler{
		articleSrv,
	}
}

// CreateArticle implements ArticleHandler.
func (h *articleHandler) CreateArticle(ctx *gin.Context) {
	var form requests.ArticleForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		errs.HandleBindingErr(ctx, err)
		return
	}
	article, err := h.articleSrv.Create(ctx, &form)
	if err != nil {
		logger.Error("error occured when creating client", logger.Err(err))
		errs.AbortWithHTTPResponse(ctx, err, nil)
		return
	}
	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "success",
		Data:    article,
	})
}

// GetArticles implements ArticleHandler.
func (h *articleHandler) GetArticles(ctx *gin.Context) {
	article, err := h.articleSrv.Get(ctx)
	if err != nil {
		errs.AbortWithHTTPResponse(ctx, err, nil)
		return
	}
	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "success",
		Data:    article,
	})
}
