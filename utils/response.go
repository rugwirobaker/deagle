package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

//APISuccessResult ...
type APISuccessResult struct {
	Status string      `json:"status"`
	Result interface{} `json:"result,omitempty"`
}

//APIErrorMessage ...
type APIErrorMessage struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

//SendSuccessResult ...
func SendSuccessResult(res http.ResponseWriter, data interface{}) {
	res.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(res)
	encoder.SetEscapeHTML(false)
	encoder.Encode(APISuccessResult{"ok", data})
}

//SendErrorResult ...
func SendErrorResult(res http.ResponseWriter, err error) {
	res.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(res)
	encoder.SetEscapeHTML(false)
	obj, ok := err.(interface{ Status() int })
	if ok == true {
		res.WriteHeader(obj.Status())
	}
	m := func(r string) string {
		if r == "" {
			return r
		}
		return strings.ToUpper(string(r[0])) + string(r[1:])
	}(err.Error())
	encoder.Encode(APIErrorMessage{"error", m})
}
