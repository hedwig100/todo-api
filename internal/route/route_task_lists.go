package route

import (
	"errors"
	"net/http"
)

// /task-lists
// POST
func taskListsCreate(writer http.ResponseWriter, request *http.Request) {

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
