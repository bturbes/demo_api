package math

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SubtractResponse holds the JSON response for calls to the Subtract API.
type SubtractResponse struct {
	A      float64 `json:"valueA"`
	B      float64 `json:"valueB"`
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Subtract some numbers
func Subtract(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subResp := SubtractResponse{}

	a, err := strconv.ParseFloat(params["a"], 64)
	if err != nil {
		subResp.Error = "First value was not a number."
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(subResp)
		return
	}

	b, err := strconv.ParseFloat(params["b"], 64)
	if err != nil {
		subResp.Error = "Second value was not a number."
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(subResp)
		return
	}
	subResp.A = a
	subResp.B = b
	subResp.Result = a - b
	json.NewEncoder(w).Encode(subResp)
}
