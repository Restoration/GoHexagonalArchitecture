package response

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"hex-go/internal/domain"
	"strconv"
)

// ArticleIndex /articles/index のレスポンス
func ArticleIndex(context *gin.Context, d []domain.Article, err error) {
	var response domain.Articles
	response.Articles = d
	if err != nil {
		Error(context, http.StatusBadRequest, 0, err)
		return
	}
	context.JSON(http.StatusOK, response)
}

// ArticleCreate /articles/create のレスポンス
func ArticleCreate(context *gin.Context, d int, err error) {
	var response domain.ArticleCommonResponse
	response.LastID = strconv.Itoa(d)
	response.Message = "success"
	if err != nil {
		Error(context, http.StatusBadRequest, 0, err)
		return
	}
	context.JSON(http.StatusOK, response)
}

// ArticleUpdate /articles/update のレスポンス
func ArticleUpdate(context *gin.Context, d bool, err error) {
	var response domain.ArticleCommonResponse
	response.Message = "success"
	if err != nil {
		Error(context, http.StatusBadRequest, 0, err)
		return
	}
	context.JSON(http.StatusOK, response)
}

// ArticleDelete /articles/delete のレスポンス
func ArticleDelete(context *gin.Context, d int, err error) {
	var response domain.ArticleCommonResponse
	response.LastID = strconv.Itoa(d)
	response.Message = "success"
	if err != nil {
		Error(context, http.StatusBadRequest, 0, err)
		return
	}
	context.JSON(http.StatusOK, response)
}
