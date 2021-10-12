package main

import (
	"net/http"
)

func server() (server http.Server) {
	mux := http.NewServeMux()
	mux.HandleFunc("/task/insert", postTask)
	mux.HandleFunc("/task/done", updateTask)
	mux.HandleFunc("/task/delete", deleteTask)
	mux.HandleFunc("/tasks/all", getAllTasks)
	mux.HandleFunc("/tasks/done", getDoneTasks)
	mux.HandleFunc("/tasks/doing", getDoingTasks)

	server = http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	return
}
