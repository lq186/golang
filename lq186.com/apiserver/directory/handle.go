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

	if dir.Lang == "" {
		dir.Lang = filterData[common.Lang].(string)
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
	data["SerNo"] = dir.SerNo
	data["Lang"] = dir.Lang

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: data})
}

func UpdateHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {
	tokenUser := filterData[common.TokenUser].(*db.User)

	var dir db.Directory
	err := common.JsonUnmarshal(request, &dir)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	err = Update(&dir, tokenUser);
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	data := make(map[string]interface{})
	data["ID"] = dir.ID
	data["Name"] = dir.DirName
	data["PID"] = dir.PID
	data["SerNo"] = dir.SerNo
	data["Lang"] = dir.Lang

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: data})
}

func ListAllHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {

	lang := filterData[common.Lang].(string)

	dirs, err := ListAll(lang)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: "Query directory error, more info: " + err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: dirs})
}

func RemoveHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {
	tokenUser := filterData[common.TokenUser].(*db.User)

	var dir db.Directory
	err := common.JsonUnmarshal(request, &dir)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	if !common.CheckEmptyParam(writer, "{ID}", dir.ID) {
		return
	}

	err = Remove(&dir, tokenUser);
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: dir.ID})
}