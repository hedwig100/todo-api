package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hedwig100/todo-api/cmd/data"
)

type UsersJson struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

// /users/
func users(writer http.ResponseWriter, request *http.Request) {
	var err error
	switch request.Method {
	case "POST":
		err = postUsers(writer, request)
	case "DELETE":
		err = deleteUsers(writer, request)
	default:
		err = errors.New("this method is not used")
	}

	if err != nil {
		sendErrorMessage(writer, err.Error())
	}
}

// /users/
// POST
func postUsers(writer http.ResponseWriter, request *http.Request) (err error) {
	len := request.ContentLength
	body := make([]byte, len)
	_, err = request.Body.Read(body)
	if err != nil {
		return
	}

	userJ := UsersJson{}
	err = json.Unmarshal(body, &userJ)
	if err != nil {
		return
	}
	user, err := data.UserCreate(userJ.Username, userJ.Password)
	if err != nil {
		return
	}

	writer.Header().Set("Location", fmt.Sprintf("/users/%s", user.Username))
	writer.WriteHeader(201)
	return
}

// /users/
// DELETE
func deleteUsers(writer http.ResponseWriter, request *http.Request) (err error) {
	len := request.ContentLength
	body := make([]byte, len)
	_, err = request.Body.Read(body)
	if err != nil {
		return
	}

	userJ := UsersJson{}
	json.Unmarshal(body, &userJ)
	_, success, err := data.Login(userJ.Username, userJ.Password)
	if err != nil || !success {
		return
	}

	err = data.UserDelete(userJ.Username)
	if err != nil {
		return
	}

	writer.WriteHeader(200)
	return
}

// /users/login
// POST
func loginUsers(writer http.ResponseWriter, request *http.Request) {
	len := request.ContentLength
	body := make([]byte, len)
	_, err := request.Body.Read(body)
	if err != nil {
		sendErrorMessage(writer, "cannot read json request")
	}

	userJ := UsersJson{}
	json.Unmarshal(body, &userJ)
	_, success, err := data.Login(userJ.Username, userJ.Password)
	if err != nil || !success {
		sendErrorMessage(writer, "password is not valid,cannot login ")
	}

	writer.WriteHeader(201)
}

type TaskListResponse struct {
	TaskLists []data.TaskList `json:"taskLists"`
}

// /users/task-lists/{username}
// GET
func getUsersTaskLists(writer http.ResponseWriter, request *http.Request) {

	// urlからusernameをだす
	username, err := isCorrectURL("/users/task-lists/", request.URL)
	if err != nil {
		sendErrorMessage(writer, err.Error())
	}

	// read body
	len := request.ContentLength
	body := make([]byte, len)
	_, err = request.Body.Read(body)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}

	// read json
	userJ := UsersJson{}
	err = json.Unmarshal(body, &userJ)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}
	password := userJ.Password

	// valid password
	_, success, err := data.Login(username, password)
	if !success || err != nil {
		sendErrorMessage(writer, "password is not valid")
		return
	}

	// get task-lists
	taskLists, err := data.UsersTaskLists(username)
	if err != nil {
		sendErrorMessage(writer, err.Error())
		return
	}
	body, _ = json.Marshal(TaskListResponse{TaskLists: taskLists})
	writer.WriteHeader(200)
	writer.Write(body)
}
