package app

import (
	"net/http"

	"github.com/hedwig100/todo-app/internal/route"
)

func GetServer() (server *http.Server) {
	mux := route.GetMux()
	server = &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}
	return
}
