package math

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// SubtractResponse holds the JSON response for calls to the Subtract API.
type SubtractResponse struct {
	A      float64 `json:"valueA"`
	B      float64 `json:"valueB"`
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Subtract some numbers
func Subtract(c echo.Context) error {
	subResp := SubtractResponse{}

	a, err := strconv.ParseFloat(c.Param("a"), 64)
	if err != nil {
		subResp.Error = "First value was not a number."
		return c.JSON(http.StatusBadRequest, subResp)
	}

	b, err := strconv.ParseFloat(c.Param("b"), 64)
	if err != nil {
		subResp.Error = "Second value was not a number."
		return c.JSON(http.StatusBadRequest, subResp)
	}
	subResp.A = a
	subResp.B = b
	subResp.Result = a - b
	return c.JSON(http.StatusOK, subResp)
}
