package router

import (
	"github.com/lq186/golang/lq186.com/apiserver/filter"
	"github.com/lq186/golang/lq186.com/apiserver/user"
	"net/http"
	"github.com/lq186/golang/lq186.com/apiserver/directory"
	"github.com/lq186/golang/lq186.com/apiserver/article"
)

func init()  {
	filter.Add("/api/.*", filter.OptionsHandle)
	filter.Add("/api/.*", filter.LangHandle)
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

	http.HandleFunc("/api/auth/art/add", filter.Filter(article.AddHandle))
	http.HandleFunc("/api/auth/art/update", filter.Filter(article.UpdateHandle))
	http.HandleFunc("/api/auth/art/remove", filter.Filter(article.RemoveHandle))
	http.HandleFunc("/api/art/list-all", filter.Filter(article.ListAllHandle))
	http.HandleFunc("/api/art/list-page", filter.Filter(article.ListPageHandle))
	http.HandleFunc("/api/art/detail", filter.Filter(article.DetailHandle))
}
