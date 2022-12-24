package query

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

// function for fetch quering the articles
func GetAllArticles(article *[]models.Article) {
	initializers.DB.Find(&article)
}