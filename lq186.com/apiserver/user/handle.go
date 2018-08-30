package user

import (
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"net/http"
	"time"
)

func LoginHandle(writer http.ResponseWriter, request *http.Request) {
	var loginBody LoginBody
	err := common.JsonUnmarshal(request, &loginBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	if !checkLoginBody(writer, &loginBody){
		return
	}

	loginBody.Ip = request.RemoteAddr

	user, err := Login(&loginBody)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.BusinessError, Message: err.Error()})
		return
	}

	data := make(map[string]interface{})
	data["Token"] = user.Token
	data["ExpirseAt"] = -time.Since(user.TokenExpirseAt) / time.Second
	data["Nickname"] = user.Nickname
	data["HeadImg"] = user.HeadImg

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: data})
}

func checkLoginBody(writer http.ResponseWriter, body *LoginBody) bool {
	return common.CheckEmptyParam(writer, "username", body.Username) && common.CheckEmptyParam(writer, "password", body.Password)
}

func AddHandle(writer http.ResponseWriter, request *http.Request) {
	var user db.User
	err := common.JsonUnmarshal(request, &user)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	err = Create(&user)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: user.ID})
}

func UpdateHandle(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	token, err := common.RequestForm(request, "token", true)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.TokenError, Message: err.Error()})
		return
	}

	var user db.User
	err = common.JsonUnmarshal(request, &user)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: err.Error()})
		return
	}

	user.Token = token

	err = Update(&user)
	if err != nil {
		response.WriteJsonData(writer, response.Data{Code: response.DBError, Message: err.Error()})
		return
	}

	data := make(map[string]interface{})
	data["Nickname"] = user.Nickname
	data["HeadImg"] = user.HeadImg

	response.WriteJsonData(writer, response.Data{Code: response.Success, Data: data})
}

type LoginBody struct {
	Username	string `json:username`
	Password	string `json:password`
	Ip			string
}
