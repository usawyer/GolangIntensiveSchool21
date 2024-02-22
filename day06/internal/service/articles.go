package service

import (
	"day06/models"
)

type Store interface {
	AddArticle(article models.Articles) error
	GetArticle(id int) (models.Articles, error)
	GetArticles(params models.SearchParams) ([]models.Articles, error)
	RemoveArticle(id int) error
}
