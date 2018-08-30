package filter

import (
	"fmt"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/user"
)

func TokenHandle(w http.ResponseWriter, r *http.Request) bool {

	err := r.ParseForm()
	if err != nil {
		log.Log.Errorf("Can not parse request form, More info: %v", err)
		response.WriteJsonData(w, response.Data{Code:response.SystemError, Message:err.Error()})
		return false
	}

	if len(r.Form["token"]) == 0 {
		response.WriteJsonData(w, response.Data{Code:response.TokenError, Message: "No token found."})
		return false
	}

	token := r.Form["token"][0]
	if "" == token {
		response.WriteJsonData(w, response.Data{Code:response.TokenError, Message: "No token found."})
		return false
	}

	if !checkedToken(token) {
		response.WriteJsonData(w, response.Data{Code:response.TokenError, Message: "Invalid token found."})
		return false
	}

	return true

}

// checkedToken check token in db and whether available
func checkedToken(token string) bool {
	fmt.Println("check token: ", token)
	return user.ExistsValidToken(token)
}