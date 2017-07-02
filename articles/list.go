package articles

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bturbes/demo_api/database"
	"github.com/bturbes/demo_api/models"
)

// List all the articles.
func List(c echo.Context) error {
	articles := []models.Article{}
	if err := database.Db.Find(&articles).Error; err != nil {
		return c.String(http.StatusInternalServerError, "There was an error creating article.")
	}
	return c.JSONPretty(http.StatusOK, &articles, "  ")
}
