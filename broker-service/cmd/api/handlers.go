package main
import (
	"net/http"
	"encoding/json"
)

type jsonResponse struct {
	Error bool  `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}
func (app *Config) Broker(w http.ResponseWriter, r *http.Request){
	payload := jsonResponse{
		Error: false,
		Message: "Broker Service",
	}
	out,_:= json.MarshalIndent(payload, "", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}