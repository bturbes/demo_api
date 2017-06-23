package math

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// AddResponse holds the JSON response for calls to the Add API.
type AddResponse struct {
	A      float64 `json:"valueA"`
	B      float64 `json:"valueB"`
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Add some numbers
func Add(c echo.Context) error {
	addResp := AddResponse{}

	a, err := strconv.ParseFloat(c.Param("a"), 64)
	if err != nil {
		addResp.Error = "First value was not a number."
		return c.JSON(http.StatusBadRequest, addResp)
	}

	b, err := strconv.ParseFloat(c.Param("b"), 64)
	if err != nil {
		addResp.Error = "Second value was not a number."
		return c.JSON(http.StatusBadRequest, addResp)
	}
	addResp.A = a
	addResp.B = b
	addResp.Result = a + b
	return c.JSON(http.StatusOK, addResp)
}
