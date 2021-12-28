package route

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/hedwig100/todo-app/internal/data"
)

var mux http.ServeMux
var writer *httptest.ResponseRecorder

// テストデータ
var createdUsername = []string{"hedwig100", "pokemon", "mac"}
var createdPassword = []string{"iajgo3o", ")8hgiau", "uhaig1928"}
var createdTaskListname = []string{"mid-term test", "for presentation", "cooking for chistmas"}
var createdTaskListId []int

// REVIEW:よりよいテストの仕方,依存性の注入?
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = GetMux()
	writer = httptest.NewRecorder()

	// データを消す
	data.Db.Exec("DELETE FROM users") // REVIEW:設計としてこういうのはよくないかも?

	// テストデータを挿入する
	for index, username := range createdUsername {
		data.UserCreate(username, createdPassword[index])
	}
	var taskList data.TaskList
	for index, listname := range createdTaskListname {
		taskList, _ = data.TaskListCreate(
			createdUsername[index],
			"oihgo3",
			listname,
		)
		createdTaskListId[index] = taskList.ListId
	}
}

func TestCreateUser(t *testing.T) {
	// 名前が被らないならユーザを作成できる
	json := strings.NewReader(`{
		"username":"fagge9092",
		"password":"98h9ghiafe"
	}`)
	request, err := http.NewRequest("POST", "/users", json)
	if err != nil {
		t.Error(err)
	}
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("cannot create user")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	// 名前が同じならユーザは作成できない
	json = strings.NewReader(fmt.Sprintf(`{
		"username": "%s", 
		"password":"9g8aho3"
	}`, createdUsername[0]))
	request, err = http.NewRequest("POST", "/users", json)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 500 {
		t.Error("can create user although this name is already used")
	}
}

func TestDeleteUser(t *testing.T) {
	// 存在するユーザがdeleteできる
	json := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[1], createdPassword[1]))
	request, _ := http.NewRequest("DELETE", "/users", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Error("cannot delete user")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	// 存在するがpasswordが違う場合はdeleteできない
	json = strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[1], createdPassword[0]))
	request, _ = http.NewRequest("DELETE", "/users", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 500 {
		t.Error("can delete user although password is not correct")
	}

	// 存在しないuserを消そうとするとエラー
	json = strings.NewReader(`{
		"username":"982ni",
		"password":"fahi"
	}`)
	request, _ = http.NewRequest("DELETE", "/users", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 500 {
		t.Error("can delete user although the user is not created")
	}
}

func TestLogin(t *testing.T) {
	// 存在するユーザにログインできること
	json := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[2], createdPassword[2]))
	request, _ := http.NewRequest("POST", "/users/login", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 201 {
		t.Error("cannot login user")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	// passwordが違うとログインできないこと
	json = strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[2], createdPassword[1]))
	request, _ = http.NewRequest("POST", "/users/login", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 500 {
		t.Fatal("can login although password is not correct")
	}

	// 存在しないユーザにログインできないこと
	json = strings.NewReader(`{
		"username":"gea",
		"password":"og0983a"
	}`)
	request, _ = http.NewRequest("POST", "/users/login", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 500 {
		t.Fatal("can login although the user is not created")
	}
}

func TestCreateTaskList(t *testing.T) {
	// 作成できること
	json := strings.NewReader(fmt.Sprintf(`{
		"username": "%s", 
		"password": "%s",
		"icon": "add",
		"listname": "textbooks I want to read"
	}`, createdUsername[0], createdPassword[0]))
	request, err := http.NewRequest("POST", "/task-lists/", json)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("cannot create task-lists")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	// パスワードが違うと作成できないこと
	json = strings.NewReader(fmt.Sprintf(`{
		"username": "%s", 
		"password": "%s",
		"icon": "add",
		"listname": "textbooks I want to read"
	}`, createdUsername[0], createdPassword[1]))
	request, err = http.NewRequest("POST", "/task-lists/", json)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 500 {
		t.Error("can create task-lists although password is not valid")
	}
}

func TestGetTaskList(t *testing.T) {
	// taskをgetできること
	request, err := http.NewRequest("GET", fmt.Sprintf("/task-lists/%d", createdTaskListId[0]), nil)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("cannot get ")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

}

func TestUpdateTaskList(t *testing.T) {
	t.Skip()
}

func TestDeleteTaskList(t *testing.T) {
	t.Skip()
}

func TestCreateTask(t *testing.T) {
	t.Skip()
}

func TestGetTask(t *testing.T) {
	t.Skip()
}

func TestUpdateTask(t *testing.T) {
	t.Skip()
}

func TestDeleteTask(t *testing.T) {
	t.Skip()
}
