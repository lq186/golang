package filter

import (
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/user"
	"github.com/lq186/golang/lq186.com/apiserver/common"
)

func TokenHandle(w http.ResponseWriter, r *http.Request, data map[string]interface{}) bool {

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

	tokenUser, err := user.TokenUser(token)
	if err != nil {
		response.WriteJsonData(w, response.Data{Code:response.TokenError, Message: err.Error()})
		return false
	}

	data[common.TokenUser] = tokenUser
	return true

}