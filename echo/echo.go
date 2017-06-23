package echo

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Echo the Request body
func Echo(c echo.Context) error {
	reader := bufio.NewReader(c.Request().Body)
	fmt.Fprintf(c.Response().Writer, "Echo: ")
	return c.Stream(http.StatusOK, "text/plain", reader)
}
