package route

import (
	"errors"
	"net/http"
)

func taskCreate(writer http.ResponseWriter, request *http.Request) {

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
