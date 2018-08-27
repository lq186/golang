package user

import (
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	response.WriteJsonData(writer, response.Data{Code:response.Success, Data:"Signin Success"})
}
