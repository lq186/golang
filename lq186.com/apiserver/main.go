package main

import (
	"flag"
	"github.com/lq186/golang/lq186.com/apiserver/config"
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"github.com/lq186/golang/lq186.com/apiserver/router"
	"net/http"
)

const (
	defaultHost = "0.0.0.0:8080"
)

func main() {
	defer db.DB.Close()

	router.AddRouter()
	listenHost := getListenHost()
	log.Log.Infof("Server start listen at: %s", listenHost)
	err := http.ListenAndServe(listenHost, nil)
	if err != nil {
		log.Log.Errorf("Can not listen and service at: %s, more info: %v", listenHost, err)
	}
}

func getListenHost() string {
	listenHost := flag.String("host", "", "Server listen host, default is http://0.0.0.0:8080")
	if "" == *listenHost {
		conf := config.Config()
		if "" == conf.Host {
			*listenHost = defaultHost
		} else {
			*listenHost = conf.Host
		}
	}
	return *listenHost
}
