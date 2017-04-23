package echo

import (
	"bufio"
	"fmt"
	"net/http"
)

// Echo the body
func Echo(w http.ResponseWriter, r *http.Request) {
	reader := bufio.NewReader(r.Body)
	fmt.Fprint(w, "Echo: ")
	reader.WriteTo(w)
}
