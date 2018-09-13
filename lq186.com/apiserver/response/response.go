package response

import (
	"encoding/json"
	"net/http"
)

const (
	Success = 0
)

const (
	SystemError = -99
	JsonError   = -1
	TokenError 	= -2
	ParamError	= -3
	DBError	= -4
	BusinessError = -5
)

func WriteJsonData(writer http.ResponseWriter, data Data) {
	addJsonHeader(writer)

	jsonData, err := json.Marshal(data)
	if err != nil {
		WriteJsonData(writer, Data{Code: JsonError, Message: err.Error()})
		return
	}

	writer.Write([]byte(jsonData))
}

func addJsonHeader(response http.ResponseWriter) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Headers", "Content-type")
	response.Header().Add("Access-Control-Allow-Headers", "Token")
	response.Header().Add("Access-Control-Allow-Headers", "Lang")
	response.Header().Add("Content-Type", "application/json;charset=UTF-8")
}

type Data struct {
	Code    int
	Message string
	Data    interface{}
}
