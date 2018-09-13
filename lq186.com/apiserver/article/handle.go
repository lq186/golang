package article

import (
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/response"
)

func AddHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {

	tokenUser := filterData[common.TokenUser].(*db.User)

	var requestBody CreateRequestBody
	err := common.JsonUnmarshal(request, &requestBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	if !common.CheckEmptyParam(writer, "DirID", requestBody.DirID) {
		return
	}

	if !common.CheckEmptyParam(writer, "Title", requestBody.Title) {
		return
	}

	if !common.CheckEmptyParam(writer, "Content", requestBody.Content) {
		return
	}

	if requestBody.Lang == "" {
		requestBody.Lang = filterData[common.Lang].(string)
	}

	article, err := Create(&requestBody, tokenUser);
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: article})
}

func UpdateHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {
	tokenUser := filterData[common.TokenUser].(*db.User)

	var requestBody UpdateRequestBody
	err := common.JsonUnmarshal(request, &requestBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	if !common.CheckEmptyParam(writer, "ID", requestBody.ID) {
		return
	}

	if !common.CheckEmptyParam(writer, "DirID", requestBody.DirID) {
		return
	}

	if !common.CheckEmptyParam(writer, "Title", requestBody.Title) {
		return
	}

	if !common.CheckEmptyParam(writer, "Content", requestBody.Content) {
		return
	}

	article, err := Update(&requestBody, tokenUser);
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: article})
}

func ListPageHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {

	var requestBody ListPageRequestBody
	err := common.JsonUnmarshal(request, &requestBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	if requestBody.Page < 1 {
		requestBody.Page = 1
	}

	if requestBody.Size < 1 {
		requestBody.Size = 10
	}

	requestBody.Lang = filterData[common.Lang].(string)

	page, err := ListPage(&requestBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: "Query directory error, more info: " + err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: page})
}

func RemoveHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {
	tokenUser := filterData[common.TokenUser].(*db.User)

	var requestBody RemoveRequestBody
	err := common.JsonUnmarshal(request, &requestBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	if requestBody.IDS == nil || len(requestBody.IDS) == 0 {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: "parameter {IDS} should not be empty."})
		return
	}

	err = Remove(&requestBody, tokenUser);
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: requestBody.IDS})
}

func DetailHandle(writer http.ResponseWriter, request *http.Request, filterData map[string]interface{}) {

	id, err := common.RequestForm(true, request, "id", true)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	detail, err := QueryDetail(id)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: detail})
}


type CreateRequestBody struct {
	DirID		string
	Title		string
	Content		string
	IsTop		bool
	Lang		string
}

type UpdateRequestBody struct {
	ID			string
	DirID		string
	Title		string
	Content		string
	IsTop		bool
	Lang		string
}

type RemoveRequestBody struct {
	IDS			[]string
}

type ListPageRequestBody struct {
	Lang 		string
	Page		uint
	Size		uint
}