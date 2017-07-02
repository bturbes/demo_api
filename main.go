package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/bturbes/demo_api/articles"
	"github.com/bturbes/demo_api/database"
	myEcho "github.com/bturbes/demo_api/echo"
	"github.com/bturbes/demo_api/formats"
	"github.com/bturbes/demo_api/math"
	"github.com/bturbes/demo_api/models"
)

func main() {
	var err error
	database.Db, err = gorm.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Db.Close()
	models.MigrateAll()

	e := echo.New()
	e.GET("/", Index)
	e.GET("/add/:a/:b", math.Add)
	e.GET("/subtract/:a/:b", math.Subtract)
	e.GET("/sampledata", formats.SampleData)
	e.POST("/echo", myEcho.Echo)
	e.POST("/articles/create", articles.Create)
	e.GET("/articles/list", articles.List)

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	log.Fatal(e.Start(":8080"))
}

// Index Welcome message
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Ben's API of sorts!")
}
