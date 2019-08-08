package http

import (
	"go-cache/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server)Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	if e := http.ListenAndServe(":12345", nil); e != nil {
		panic(e)
	}
}

func (s *Server)cacheHandler() http.Handler {
    return &cacheHandler{s}
}

func (s *Server)statusHandler() http.Handler {
	return &statusHandler{s}
}
