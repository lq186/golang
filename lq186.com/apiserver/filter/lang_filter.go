package filter

import (
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"net/http"
)

func LangHandle(w http.ResponseWriter, r *http.Request, data map[string]interface{}) bool {

	lang := r.Header.Get("Lang")
	if "" == lang {
		response.WriteJsonData(w, response.Data{Code:response.TokenError, Message: "No Lang found."})
		return false
	}

	data[common.Lang] = lang
	return true

}