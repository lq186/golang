package directory

import (
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/response"
)

func AddHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {

	tokenUser := filterData[common.TokenUser].(*db.User)

	var dir db.Directory
	err := common.JsonUnmarshal(request, &dir)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	err = Create(&dir, tokenUser);
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	data := make(map[string]interface{})
	data["ID"] = dir.ID
	data["Name"] = dir.DirName
	data["PID"] = dir.PID

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: data})
}
