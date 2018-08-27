package router

import (
	"github.com/lq186/golang/lq186.com/apiserver/handle/user"
	"net/http"
)

func AddRouter()  {
	http.HandleFunc("/api/user/login", user.Login)
}
