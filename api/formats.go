package api

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sampleStruct struct {
	XMLName xml.Name `json:"-" xml:"sampleData"`
	ID      int      `json:"id" xml:"id"`
	Value   string   `json:"value" xml:"value"`
	List    []string `json:"list" xml:"list"`
}

// SampleData renders
func SampleData(c *gin.Context) {
	format := c.Query("format")
	sample := sampleStruct{ID: 1, Value: "I hope this works.", List: []string{"test", "123", "lol"}}

	if format == "xml" {
		c.XML(http.StatusOK, sample)
		return
	}
	c.JSON(http.StatusOK, sample)
}
