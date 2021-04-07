package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"hex-go/internal/datasource"
	"hex-go/internal/domain"
)

// ArticleRepository .
type ArticleRepository struct {
	DB *gorm.DB
}

// NewArticleRepository .
func NewArticleRepository() *ArticleRepository {
	repo := new(ArticleRepository)
	repo.DB = datasource.DBConn()
	return repo
}

// FindArticlel articleを取得する
func (a *ArticleRepository) FindArticle(offset string, limit string) ([]domain.Article, error) {
	var articles []domain.Article
	a.DB.Offset(offset).Limit(limit).Find(&articles)
	return articles, nil
}

// Create 保存処理
func (a *ArticleRepository) Create(params domain.ArticleParams) (int, error) {
	d := domain.Article{Title: params.Title, Body: params.Body}
	err := a.DB.Table("articles").Save(&d).Error
	if err != nil {
		return 0, err
	}
	return d.ID, nil
}

// Update 更新処理
func (a *ArticleRepository) Update(params domain.ArticleParams) (bool, error) {
	id := params.ID
	d := domain.Article{Title: params.Title, Body: params.Body, UpdatedAt: time.Now()}
	a.DB.Table("articles").Where("id = ?", id).Update(&d)
	return true, nil
}

// Delete 削除処理
func (a *ArticleRepository) Delete(id int) (int, error) {
	var article domain.Article
	err := a.DB.Table("articles").Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return article.ID, err
	}
	return article.ID, nil
}
