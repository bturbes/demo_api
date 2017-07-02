package articles

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/bturbes/demo_api/database"
	"github.com/bturbes/demo_api/models"
)

// Create a new article.
func Create(c echo.Context) error {
	a := new(models.Article)
	if err := c.Bind(a); err != nil {
		return err
	}
	database.Db.Create(&a)
	return c.String(http.StatusOK, "Created the article")
}
