package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rtanx/gostarter/internal/http/dto/responses"
	"github.com/rtanx/gostarter/internal/http/handlers"
	"github.com/rtanx/gostarter/internal/http/middlewares"
	"github.com/rtanx/gostarter/internal/repositories"
	"github.com/rtanx/gostarter/internal/services"
)

func RoutesHandler(r *gin.Engine) {

	articleRepo := repositories.NewArticleRepository()
	articleSrvc := services.NewArticleService(articleRepo)
	articleHndlr := handlers.NewArticleHandler(articleSrvc)

	r.Use(middlewares.EnableCors())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Pong",
		})
	})

	v1 := r.Group("/v1")
	{
		article := v1.Group("/article")
		{
			article.GET("", articleHndlr.GetArticles)
		}
	}
}
