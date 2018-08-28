package filter

import (
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"net/http"
	"regexp"
)

var (
	filters = []*filterHandle{}
)

func Add(pattern string, filter Handle) {
	filters = append(filters, &filterHandle{pattern:pattern,filter:filter})
}

func Filter(handle WebHandle) WebHandle {
	
	return func(w http.ResponseWriter, r *http.Request) {

		for _, filter := range filters {
			uri := r.RequestURI + "/"
			matched, err := regexp.MatchString(filter.pattern, uri)
			if err != nil || !matched {
				continue
			}

			if !filter.filter(w, r) {
				log.Log.Debugf("filter: %v not passed.", filter)
				return
			}
		}

		handle(w, r)
	}
	
}

type filterHandle struct {
	pattern string
	filter Handle
}

type Handle func(w http.ResponseWriter, r *http.Request) bool

type WebHandle func(w http.ResponseWriter, r *http.Request)