package route

import (
	"encoding/json"
	"errors"
	"net/http"

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

// /task-lists/{listId}
// GET
func taskListsGet(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}

// /task-lists/{listId}
// PUT
func taskListsUpdate(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}

// /task-lists/{listId}
// DELETE
func taskListsDelete(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}
