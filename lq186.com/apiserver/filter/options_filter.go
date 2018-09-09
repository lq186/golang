package filter

import (
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"net/http"
)

func OptionsHandle(w http.ResponseWriter, r *http.Request, data map[string]interface{}) bool {

	if "OPTIONS" == r.Method {
		w.Header().Add("Access-Control-Max-Age", "3600");
		response.WriteJsonData(w, response.Data{Code:response.Success})
		return false
	}

	return true

}