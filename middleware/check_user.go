package middleware

import (
	"fmt"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CheckUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId")
	if userId == nil {
        http.Error(writer, "Unauthorized", http.StatusUnauthorized)
        return
	}

    userID := userId.(int) 
    fmt.Printf("User ID from context: %d\n", userID)
	hello := "hello user " + strconv.Itoa(userID)

	helper.WriteEncodeResponse(writer, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: hello,
	})


}
