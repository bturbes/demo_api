package formats

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type sampleStruct struct {
	XMLName xml.Name `json:"-" xml:"sampleData"`
	ID      int      `json:"id" xml:"id"`
	Value   string   `json:"value" xml:"value"`
	List    []string `json:"list" xml:"list"`
}

// SampleData renders
func SampleData(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	sample := sampleStruct{ID: 1, Value: "I hope this works.", List: []string{"test", "123", "lol"}}

	if format == "xml" {
		xml.NewEncoder(w).Encode(sample)
	} else {
		json.NewEncoder(w).Encode(sample)
	}

}
