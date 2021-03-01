package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DecodePathVariable(val string, r *http.Request)string  {
	param := mux.Vars(r)
	return param[val]
}
func JsonDecoder(val interface{}, r *http.Request) error {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&val)
	if err != nil {
		return err
	}
	return nil
}
