package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/bturbes/demo_api/api"
	"github.com/bturbes/demo_api/database"
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

	r := gin.Default()
	r.GET("/", Index)

	ag := r.Group("/api")
	ag.GET("/add/:a/:b", api.Add)
	ag.GET("/subtract/:a/:b", api.Subtract)
	ag.POST("/echo", api.Echo)
	ag.GET("/sampledata", api.SampleData)
	ag.POST("/articles/create", api.CreateArticle)
	ag.GET("/articles/list", api.ListArticles)

	log.Fatal(r.Run(":8080"))
}

// Index Welcome message
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Ben's API of sorts!")
}
