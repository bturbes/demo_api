package api

import (
	"net/http"

	"github.com/bturbes/demo_api/database"
	"github.com/bturbes/demo_api/models"
	"github.com/gin-gonic/gin"
)

// ListArticles lists all the articles.
func ListArticles(c *gin.Context) {
	articles := []models.Article{}
	if err := database.Db.Find(&articles).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error listing articles")
		return
	}
	c.JSON(http.StatusOK, &articles)
}

// CreateArticle creates a new article.
func CreateArticle(c *gin.Context) {
	a := new(models.Article)
	if err := c.BindJSON(a); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}
	if err := database.Db.Create(&a).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error creating article")
		return
	}
	c.String(http.StatusOK, "Article created")
}
