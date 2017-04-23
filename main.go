package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/bturbes/demo_api/echo"
	"github.com/bturbes/demo_api/formats"
	"github.com/bturbes/demo_api/math"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Index)
	r.HandleFunc("/add/{a}/{b}", math.Add)
	r.HandleFunc("/subtract/{a}/{b}", math.Subtract)
	r.HandleFunc("/echo", echo.Echo)
	r.HandleFunc("/sampledata", formats.SampleData)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(UserAgent))
	n.UseHandler(http.HandlerFunc(AllowedOrigins))
	n.UseHandler(r)

	log.Println("Listening on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", n))
}

// Index Welcome message
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Ben's API of sorts!")
}

// AllowedOrigins sets Access-Control-Allow-Origin
func AllowedOrigins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// UserAgent logs
func UserAgent(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userAgent := r.Header.Get("User-Agent")
	log.Println("User-Agent:", userAgent)
	next(rw, r)
}
