package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/hedwig100/todo-app/internal/data"
)

type TcReq struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	ListId   int       `json:"listId"`
	Taskname string    `json:"taskname"`
	Deadline time.Time `json:"deadline"`
}

type TcRes struct {
	TaskId int `json:"taskId"`
}

func taskCreate(writer http.ResponseWriter, request *http.Request) {
	// read json
	len := request.ContentLength
	body := make([]byte, len)
	_, err := request.Body.Read(body)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}

	// valid password
	var tcR TcReq
	err = json.Unmarshal(body, &tcR)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}
	_, success, err := data.Login(tcR.Username, tcR.Password)
	if !success || err != nil {
		sendErrorMessage(writer, "password is not valid")
		return
	}

	// insert task
	taskList, err := data.TaskCreate(tcR.Username, tcR.ListId, tcR.Taskname, tcR.Deadline)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}

	body, _ = json.Marshal(TcRes{TaskId: taskList.TaskId})
	writer.WriteHeader(201)
	writer.Write(body)
}

func taskHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	switch request.Method {
	case "GET":
		err = taskGet(writer, request)
	case "PUT":
		err = taskUpdate(writer, request)
	case "DELETE":
		err = taskDelete(writer, request)
	default:
		err = errors.New("this method is not used")
	}
	if err != nil {
		sendErrorMessage(writer, err.Error())
	}
}

func taskGet(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}

func taskUpdate(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}

func taskDelete(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}
