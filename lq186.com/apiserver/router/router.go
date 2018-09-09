package router

import (
	"github.com/lq186/golang/lq186.com/apiserver/filter"
	"github.com/lq186/golang/lq186.com/apiserver/user"
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/directory"
)

func init()  {
	filter.Add("/api/.*", filter.OptionsHandle)
	filter.Add("/api/auth/.*", filter.TokenHandle)
}

func AddRouter() {
	http.HandleFunc("/api/user/login", filter.Filter(user.LoginHandle))
	http.HandleFunc("/api/user/add", filter.Filter(user.AddHandle))
	http.HandleFunc("/api/auth/user/update", filter.Filter(user.UpdateHandle))

	http.HandleFunc("/api/auth/dir/add", filter.Filter(directory.AddHandle))
	http.HandleFunc("/api/auth/dir/update", filter.Filter(directory.UpdateHandle))
	http.HandleFunc("/api/auth/dir/remove", filter.Filter(directory.RemoveHandle))
	http.HandleFunc("/api/dir/list-all", filter.Filter(directory.ListAllHandle))
}
