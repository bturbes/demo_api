package api

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Echo the Request body
func Echo(c *gin.Context) {
	reader := bufio.NewReader(c.Request.Body)
	c.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(c.Writer, "Echo: ")
	reader.WriteTo(c.Writer)
}
