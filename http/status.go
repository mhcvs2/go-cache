package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if b, e := json.Marshal(h.GetStat()); e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}  else {
		w.Write(b)
	}
}



