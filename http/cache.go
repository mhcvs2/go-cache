package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// "/cache/testkey"
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := r.Method
	if m == http.MethodPut {
		b, _ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			if e := h.Set(key, b); e != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)
			}

		}
		return
	}

	if m == http.MethodGet{
		if b, e := h.Get(key); e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		} else if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write(b)
		}
		return
	}

	if m == http.MethodDelete {
		if e := h.Del(key); e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}



