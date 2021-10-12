package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"todo-app/data"
)

var testserver http.Server
var writer *httptest.ResponseRecorder
var id int

func TestMain(m *testing.M) {
	setup()
	id = 33
	code := m.Run()
	os.Exit(code)
}

func setup() {
	testserver = server()
	writer = httptest.NewRecorder()
}

func TestPostTask(t *testing.T) {
	json := strings.NewReader(`{
		"taskname":"buy a potate",
		"deadline":"2021/10/10 21:31:02"
	}`)
	request, _ := http.NewRequest("POST", "/task/insert", json)
	testserver.Handler.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	request, _ := http.NewRequest("UPDATE", fmt.Sprintf("/task/done?id=%v", id), nil)
	testserver.Handler.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/task/delete?id=%v", id), nil)
	testserver.Handler.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestGetAllTasks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tasks/all", nil)
	testserver.Handler.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var task []data.Task
	json.Unmarshal(writer.Body.Bytes(), &task)
	for i := range task {
		log.Println(task[i].Id, task[i].Deadline, task[i].IsDone, task[i].DoneTime)
	}
}

func TestGetDoneTasks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tasks/done", nil)
	testserver.Handler.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var task []data.Task
	json.Unmarshal(writer.Body.Bytes(), &task)
	for i := range task {
		if task[i].IsDone != false {
			t.Error("Don't get done tasks")
		}
	}
}

func TestGetDoingTasks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tasks/doing", nil)
	testserver.Handler.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var task []data.Task
	json.Unmarshal(writer.Body.Bytes(), &task)
	for i := range task {
		if task[i].IsDone != true {
			t.Error("Don't get doing tasks")
		}
	}
}
