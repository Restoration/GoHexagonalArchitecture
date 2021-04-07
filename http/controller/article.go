package controller

import (
	"hex-go/http/response"
	"hex-go/internal/domain"
	"hex-go/internal/repository"
	"hex-go/internal/usecase"

	"github.com/gin-gonic/gin"
	"strconv"
)

// ArticleIndex .
func ArticleIndex(c *gin.Context) {
	var d domain.Pager
	d.Offset = c.Query("offset")
	d.Limit = c.Query("limit")
	uc := usecase.NewArticleUseCase(repository.NewArticleRepository())
	articles, err := uc.FindArticle(d)
	response.ArticleIndex(c, articles, err)
}

// ArticleCreate .
func ArticleCreate(c *gin.Context) {
	var d domain.ArticleParams
	d.Title = c.Query("title")
	d.Body = c.Query("body")
	uc := usecase.NewArticleUseCase(repository.NewArticleRepository())
	articles, err := uc.Create(d)
	response.ArticleCreate(c, articles, err)
}

// ArticleUpdate .
func ArticleUpdate(c *gin.Context) {
	var d domain.ArticleParams
	id, _ := strconv.Atoi(c.Query("id"))
	d.ID = id
	d.Title = c.Query("title")
	d.Body = c.Query("body")
	uc := usecase.NewArticleUseCase(repository.NewArticleRepository())
	articles, err := uc.Update(d)
	response.ArticleUpdate(c, articles, err)
}

// ArticleDelete .
func ArticleDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	uc := usecase.NewArticleUseCase(repository.NewArticleRepository())
	articles, err := uc.Delete(id)
	response.ArticleDelete(c, articles, err)
}
