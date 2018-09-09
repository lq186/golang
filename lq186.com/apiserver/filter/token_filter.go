package filter

import (
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"github.com/lq186/golang/lq186.com/apiserver/user"
	"net/http"
)

func TokenHandle(w http.ResponseWriter, r *http.Request, data map[string]interface{}) bool {

	token := r.Header.Get("Token")
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