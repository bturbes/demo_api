package math

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AddResponse holds the JSON response for calls to the Add API.
type AddResponse struct {
	A      float64 `json:"valueA"`
	B      float64 `json:"valueB"`
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Add some numbers
func Add(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	addResp := AddResponse{}

	a, err := strconv.ParseFloat(params["a"], 64)
	if err != nil {
		addResp.Error = "First value was not a number."
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(addResp)
		return
	}

	b, err := strconv.ParseFloat(params["b"], 64)
	if err != nil {
		addResp.Error = "Second value was not a number."
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(addResp)
		return
	}
	addResp.A = a
	addResp.B = b
	addResp.Result = a + b
	json.NewEncoder(w).Encode(addResp)
}
