package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hedwig100/todo-app/internal/data"
)

type UsersJson struct {
	Username string `json:"username"`
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
