package port

import (
	"hex-go/internal/domain"
)

// ArticlePort is
type ArticlePort interface {
	FindArticle(offset string, limit string) ([]domain.Article, error)
	Create(domain.ArticleParams) (int, error)
	Update(domain.ArticleParams) (bool, error)
	Delete(id int) (int, error)
}
