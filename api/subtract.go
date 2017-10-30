package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SubtractResponse holds the JSON response for calls to the Subtract API.
type SubtractResponse struct {
	A      float64 `json:"valueA"`
	B      float64 `json:"valueB"`
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Subtract some numbers
func Subtract(c *gin.Context) {
	subResp := SubtractResponse{}

	a, err := strconv.ParseFloat(c.Param("a"), 64)
	if err != nil {
		subResp.Error = "First value was not a number."
		c.JSON(http.StatusBadRequest, subResp)
		return
	}

	b, err := strconv.ParseFloat(c.Param("b"), 64)
	if err != nil {
		subResp.Error = "Second value was not a number."
		c.JSON(http.StatusBadRequest, subResp)
		return
	}
	subResp.A = a
	subResp.B = b
	subResp.Result = a - b
	c.JSON(http.StatusOK, subResp)
}
