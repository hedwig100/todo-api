package route

import (
	"net/http"
)

func GetMux() (mux http.ServeMux) {
	mux.HandleFunc("/users", users)
	mux.HandleFunc("/users/login", loginUsers)
	mux.HandleFunc("/users/task-lists/", getUsersTaskLists)

	mux.HandleFunc("/task-lists", taskListsCreate)
	mux.HandleFunc("/task-lists/", taskListsHandler)

	mux.HandleFunc("/tasks", taskCreate)
	mux.HandleFunc("/tasks/", taskHandler)

	return mux
}
