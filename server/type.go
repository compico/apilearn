package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Port string
	R    *httprouter.Router
	S    *http.Server
}
