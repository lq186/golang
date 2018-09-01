package router

import (
	"github.com/lq186/golang/lq186.com/apiserver/filter"
	"github.com/lq186/golang/lq186.com/apiserver/user"
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/directory"
)

func init()  {
	filter.Add("/api/.*", filter.TokenHandle)
}

func AddRouter() {
	http.HandleFunc("/api/user/login", user.LoginHandle)
	http.HandleFunc("/api/user/add", user.AddHandle)
	http.HandleFunc("/api/user/update", filter.Filter(user.UpdateHandle))

	http.HandleFunc("/api/dir/add", filter.Filter(directory.AddHandle))
	http.HandleFunc("/api/dir/update", filter.Filter(directory.UpdateHandle))
	http.HandleFunc("/api/dir/remove", filter.Filter(directory.RemoveHandle))
	http.HandleFunc("/api/dir/list-all", filter.Filter(directory.ListAllHandle))
}
