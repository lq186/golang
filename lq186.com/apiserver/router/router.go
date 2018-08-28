package router

import (
	"github.com/lq186/golang/lq186.com/apiserver/filter"
	"github.com/lq186/golang/lq186.com/apiserver/handle/user"
	"net/http"
)

func init()  {
	filter.Add("/api/user/.*", filter.TokenHandle)
}

func AddRouter() {
	http.HandleFunc("/api/user/login", filter.Filter(user.Login))
}
