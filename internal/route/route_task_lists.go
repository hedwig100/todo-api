package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/hedwig100/todo-app/internal/data"
)

// REVIEW:json.Marshalなどを使わずにjson化する
type taskListRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Icon     string `json:"icon"`
	Listname string `json:"listname"`
}

type ListIdResponse struct {
	ListId int `json:"listId"`
}

// /task-lists
// POST
func taskListsCreate(writer http.ResponseWriter, request *http.Request) {
	// read json
	len := request.ContentLength
	body := make([]byte, len)
	_, err := request.Body.Read(body)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}

	// valid password
	var requestJson taskListRequest
	json.Unmarshal(body, &requestJson)
	_, success, err := data.Login(requestJson.Username, requestJson.Password)
	if !success || err != nil {
		sendErrorMessage(writer, "password is not valid")
		return
	}

	// insert tasklist
	taskList, err := data.TaskListCreate(requestJson.Username, requestJson.Icon, requestJson.Listname)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}

	body, _ = json.Marshal(ListIdResponse{ListId: taskList.ListId})
	writer.WriteHeader(201)
	writer.Write(body)
}

func taskListsHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	switch request.Method {
	case "GET":
		err = taskListsGet(writer, request)
	case "PUT":
		err = taskListsUpdate(writer, request)
	case "DELETE":
		err = taskListsDelete(writer, request)
	default:
		err = errors.New("this method is not used")
	}
	if err != nil {
		sendErrorMessage(writer, err.Error())
	}
}

type PwRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type taskListResponse struct {
	Username string      `json:"username"`
	Icon     string      `json:"icon"`
	Listname string      `json:"listname"`
	Tasks    []data.Task `json:"tasks"`
}

// /task-lists/{listId}
// GET
func taskListsGet(writer http.ResponseWriter, request *http.Request) (err error) {
	// check if correct url is passed
	traling, err := isCorrectURL("/task-lists/", request.URL)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(traling)
	if err != nil {
		return
	}

	// parse json and login
	len := request.ContentLength
	body := make([]byte, len)
	_, err = request.Body.Read(body)
	if err != nil {
		return
	}

	var pwR PwRequest
	err = json.Unmarshal(body, &pwR)
	if err != nil {
		return
	}
	if _, success, err := data.Login(pwR.Username, pwR.Password); !success || err != nil {
		return errors.New("username and password is not valid")
	}

	// get tasklist
	tasklist, tasks, err := data.TaskListAndTasks(listId)
	if err != nil {
		return
	}

	// return response
	body, _ = json.Marshal(taskListResponse{
		Username: tasklist.Username,
		Icon:     tasklist.Icon,
		Listname: tasklist.Listname,
		Tasks:    tasks,
	})
	writer.WriteHeader(200)
	writer.Write(body)
	return
}

type LuRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Icon     string `json:"icon"`
	Listname string `json:"listname"`
}

// /task-lists/{listId}
// PUT
func taskListsUpdate(writer http.ResponseWriter, request *http.Request) (err error) {
	// check if correct url is passed
	traling, err := isCorrectURL("/task-lists/", request.URL)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(traling)
	if err != nil {
		return
	}

	// parse json and login
	len := request.ContentLength
	body := make([]byte, len)
	_, err = request.Body.Read(body)
	if err != nil {
		return
	}

	var luR LuRequest
	err = json.Unmarshal(body, &luR)
	if err != nil {
		return
	}
	if _, success, err := data.Login(luR.Username, luR.Password); !success || err != nil {
		return errors.New("username and password is not valid")
	}

	// update
	err = data.TaskListUpdate(data.TaskList{
		ListId:   listId,
		Username: luR.Username,
		Icon:     luR.Icon,
		Listname: luR.Listname,
	})
	if err != nil {
		return
	}

	// write response
	writer.WriteHeader(201)
	return
}

// /task-lists/{listId}
// DELETE
func taskListsDelete(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}
