package usecase

import (
	"hex-go/internal/domain"
	"hex-go/internal/port"
)

// ArticleUseCase is
type ArticleUseCase struct {
	articlePort port.ArticlePort
}

// NewArticleUseCase is
func NewArticleUseCase(articlePort port.ArticlePort) *ArticleUseCase {
	return &ArticleUseCase{articlePort: articlePort}
}

// FindArticle is
func (a *ArticleUseCase) FindArticle(d domain.Pager) ([]domain.Article, error) {
	return a.articlePort.FindArticle(d.Offset, d.Limit)
}

// Create .
func (a *ArticleUseCase) Create(params domain.ArticleParams) (int, error) {
	return a.articlePort.Create(params)
}

// Update .
func (a *ArticleUseCase) Update(params domain.ArticleParams) (bool, error) {
	return a.articlePort.Update(params)
}

// Delete .
func (a *ArticleUseCase) Delete(id int) (int, error) {
	return a.articlePort.Delete(id)
}
