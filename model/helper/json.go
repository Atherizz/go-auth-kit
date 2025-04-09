package helper

import (
	"golang-restful-api/model/web"
	"encoding/json"
	"net/http"
)

func WriteEncodeResponse(writer http.ResponseWriter, webResponse web.WebResponse) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicError(err)
}