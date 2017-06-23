package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	myEcho "github.com/bturbes/demo_api/echo"
	"github.com/bturbes/demo_api/formats"
	"github.com/bturbes/demo_api/math"
)

func main() {
	e := echo.New()
	e.GET("/", Index)
	e.GET("/add/:a/:b", math.Add)
	e.GET("/subtract/:a/:b", math.Subtract)
	e.GET("/sampledata", formats.SampleData)
	e.POST("/echo", myEcho.Echo)

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	log.Fatal(e.Start(":8080"))
}

// Index Welcome message
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Ben's API of sorts!")
}

// AllowedOrigins sets Access-Control-Allow-Origin
func AllowedOrigins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
