package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/data"
)

func setCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
}

// POST /task/insert/
func postTask(w http.ResponseWriter, r *http.Request) {
	setCORS(&w)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var task data.Task
	json.Unmarshal(body, &task)

	task.IsDone = false
	task.DoneTime = data.TimeZero
	err := task.Insert()
	if err != nil {
		danger(err, "Cannot insert task")
	}
	w.WriteHeader(200)
}

// UPDATE /task/done
func updateTask(w http.ResponseWriter, r *http.Request) {
	setCORS(&w)

	r.ParseForm()
	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		danger(err, "Cannot update task")
	}
	err = data.DoneTask(id)
	if err != nil {
		danger(err, "Cannot update task")
	}
	w.WriteHeader(200)
}

// DELETE /task/delete
func deleteTask(w http.ResponseWriter, r *http.Request) {
	setCORS(&w)

	r.ParseForm()
	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		danger(err, "Cannot delete task")
	}
	err = data.DeleteTask(id)
	if err != nil {
		danger(err, "Cannot delete task")
	}
	w.WriteHeader(200)
}

// GET /tasks/all
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	setCORS(&w)

	tasks, err := data.GetAllTask()
	if err != nil {
		danger(err, "Cannot get all tasks")
	}
	output, err := json.MarshalIndent(&tasks, "", "\t\t")
	if err != nil {
		danger(err, "Cannot get all tasks")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// GET /tasks/done
func getDoneTasks(w http.ResponseWriter, r *http.Request) {
	setCORS(&w)

	tasks, err := data.GetDoneTask()
	if err != nil {
		danger(err, "Cannot get done tasks")
	}
	output, err := json.MarshalIndent(&tasks, "", "\t\t")
	if err != nil {
		danger(err, "Cannot get done tasks")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// GET /tasks/dogin
func getDoingTasks(w http.ResponseWriter, r *http.Request) {
	setCORS(&w)

	tasks, err := data.GetDoingTask()
	if err != nil {
		danger(err, "Cannot get doing tasks")
	}
	output, err := json.MarshalIndent(&tasks, "", "\t\t")
	if err != nil {
		danger(err, "Cannot get doing tasks")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
