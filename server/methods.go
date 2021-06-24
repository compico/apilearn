package server

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewConfig() *Server {
	x := new(Server)
	return x
}

func (conf *Server) InitServer() {
	conf.S = &http.Server{
		Addr:           conf.Port,
		Handler:        conf.R,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (conf *Server) InitRouter() {
	conf.R = httprouter.New()
}
