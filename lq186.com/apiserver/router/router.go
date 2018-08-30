package router

import (
	"github.com/lq186/golang/lq186.com/apiserver/filter"
	"github.com/lq186/golang/lq186.com/apiserver/user"
	"net/http"
)

func init()  {
	filter.Add("/api/user/.*", filter.TokenHandle)
}

func AddRouter() {
	http.HandleFunc("/api/user/login", user.LoginHandle)
	http.HandleFunc("/api/user/add", user.AddHandle)
	http.HandleFunc("/api/user/update", filter.Filter(user.LoginHandle))
}
