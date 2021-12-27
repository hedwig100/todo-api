package route

import (
	"net/http"
)

func GetMux() (mux http.ServeMux) {
	mux.HandleFunc("/users", users)
	mux.HandleFunc("/users/login", loginUsers)

	return mux
}
