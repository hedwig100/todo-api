package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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
	// check if correct url is passed
	traling, err := isCorrectURL("/tasks/", request.URL)
	if err != nil {
		return
	}
	taskId, err := strconv.Atoi(traling)
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

	// get task
	task, err := data.TaskRetrieve(taskId)
	if err != nil {
		return
	}

	// return response
	body, _ = json.Marshal(task)
	writer.WriteHeader(200)
	writer.Write(body)
	return
}

func taskUpdate(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}

func taskDelete(writer http.ResponseWriter, request *http.Request) (err error) {
	return
}
