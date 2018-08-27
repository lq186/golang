package main

import (
	"flag"
	"github.com/lq186/golang/lq186.com/apiserver/router"
	"net/http"
)

func main() {
	listenHost := flag.String("host", "0.0.0.0:8080", "Server listen host, default is http://0.0.0.0:8080")
	router.AddRouter()
	http.ListenAndServe(*listenHost, nil)
}
