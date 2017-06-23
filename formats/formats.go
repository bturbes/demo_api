package formats

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
)

type sampleStruct struct {
	XMLName xml.Name `json:"-" xml:"sampleData"`
	ID      int      `json:"id" xml:"id"`
	Value   string   `json:"value" xml:"value"`
	List    []string `json:"list" xml:"list"`
}

// SampleData renders
func SampleData(c echo.Context) error {
	format := c.QueryParam("format")
	sample := sampleStruct{ID: 1, Value: "I hope this works.", List: []string{"test", "123", "lol"}}

	if format == "xml" {
		return c.XML(http.StatusOK, sample)
	}
	return c.JSON(http.StatusOK, sample)
}
