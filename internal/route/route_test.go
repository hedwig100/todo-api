package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/hedwig100/todo-app/internal/data"
)

var mux http.ServeMux
var writer *httptest.ResponseRecorder

// テストデータ
var createdUsername = []string{"hedwig100", "pokemon", "mac"}
var createdPassword = []string{"iajgo3o", ")8hgiau", "uhaig1928"}
var createdTaskListname = []string{"cooking for chistmas", "for presentation", "mid-term test"}
var createdTaskListId []int
var createdTaskname = []string{"buy a chicken", "buy a present for children", "buy wine"}
var createdDeadline = []time.Time{
	time.Date(2022, time.November, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2022, time.December, 24, 0, 0, 0, 0, time.UTC),
	time.Date(2022, time.December, 20, 0, 0, 0, 0, time.UTC),
}
var createdTaskId []int

// REVIEW:よりよいテストの仕方,依存性の注入?
func TestMain(m *testing.M) {
	setUp()
	m.Run()
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
	createdTaskListId = make([]int, 3)
	for index, listname := range createdTaskListname {
		taskList, _ = data.TaskListCreate(
			createdUsername[0],
			"oihgo3",
			listname,
		)
		createdTaskListId[index] = taskList.ListId
	}

	createdTaskId = make([]int, 3)
	for index, taskname := range createdTaskname {
		task, err := data.TaskCreate(
			createdUsername[0],
			createdTaskListId[0],
			taskname,
			createdDeadline[index],
		)
		if err != nil {
			log.Fatal(err.Error())
		}
		createdTaskId[index] = task.TaskId
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

func TestGetUsersList(t *testing.T) {
	jsonA := strings.NewReader(fmt.Sprintf(`{
		"password":"%s"
	}`, createdPassword[0]))
	request, _ := http.NewRequest("GET", fmt.Sprintf("/users/task-lists/%s", createdUsername[0]), jsonA)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		errMsg := writer.Body.String()
		t.Error(errMsg)
		t.Fatal("cannot get user's tasklists")
	}

	var jsonRes TaskListResponse
	err := json.Unmarshal(writer.Body.Bytes(), &jsonRes)
	if err != nil {
		t.Error(err)
	}

	sort.Slice(jsonRes.TaskLists, func(i, j int) bool { return jsonRes.TaskLists[i].Listname < jsonRes.TaskLists[j].Listname })
	for index, tasklist := range jsonRes.TaskLists {
		if tasklist.ListId != createdTaskListId[index] ||
			tasklist.Listname != createdTaskListname[index] {
			t.Fatal("cannot get correct tasklist")
		}
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
	request, err := http.NewRequest("POST", "/task-lists", json)
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
	request, err = http.NewRequest("POST", "/task-lists", json)
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
	json_ := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[0], createdPassword[0]))
	request, err := http.NewRequest("GET", fmt.Sprintf("/task-lists/%d", createdTaskListId[0]), json_)
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

	var res taskListResponse
	err = json.Unmarshal(writer.Body.Bytes(), &res)
	if err != nil {
		t.Error(err)
	}

	if res.Username != createdUsername[0] ||
		res.Icon != "oihgo3" ||
		res.Listname != createdTaskListname[0] ||
		len(res.Tasks) != len(createdTaskname) {
		t.Error("cannot get right task list")
	}

	for index, task := range res.Tasks {
		if createdTaskname[index] != task.Taskname ||
			!createdDeadline[index].Equal(task.Deadline) ||
			createdTaskId[index] != task.TaskId ||
			task.IsDone ||
			task.IsImportant ||
			task.Memo != "" {
			t.Error("cannot get right task")
		}
	}
}

func TestUpdateTaskList(t *testing.T) {
	json_ := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s",
		"icon":"%s",
		"listname":"%s"
	}`, createdUsername[0], createdPassword[0], "sub", "new listname"))
	request, err := http.NewRequest("PUT", fmt.Sprintf("/task-lists/%d", createdTaskListId[1]), json_)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("cannot update ")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	tasklist, err := data.TaskListRetrieve(createdTaskListId[1])
	if err != nil || tasklist.Icon != "sub" || tasklist.Listname != "new listname" {
		t.Error("tasklist is not updated")
	}
}

func TestDeleteTaskList(t *testing.T) {
	json_ := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[0], createdPassword[0]))
	request, err := http.NewRequest("DELETE", fmt.Sprintf("/task-lists/%d", createdTaskListId[1]), json_)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("cannot delete ")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	_, err = data.TaskListRetrieve(createdTaskListId[1])
	if err == nil {
		t.Error("data hasn't be deleted")
	}
	t.Logf("%v\n", err.Error())
}

type Response1 struct {
	TaskId int `json:"taskId"`
}

func TestCreateTask(t *testing.T) {
	// タスクが作成できる。
	taskname := "task1"
	deadline := time.Date(2021, time.November, 10, 0, 0, 0, 0, time.UTC)
	json_ := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s",
		"listId":%d,
		"taskname":"%s",
		"deadline":"%s"
	}`, createdUsername[0], createdPassword[0], createdTaskListId[0], taskname, deadline.Format(time.RFC3339)))
	request, err := http.NewRequest("POST", "/tasks", json_)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("task has not been created")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	var res Response1
	err = json.Unmarshal(writer.Body.Bytes(), &res)
	if err != nil {
		return
	}

	// listIdが異なるタスクが作成できない。
	taskname = "task1"
	deadline = time.Date(2021, time.November, 10, 0, 0, 0, 0, time.UTC)
	json_ = strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s",
		"listId":19834,
		"taskname":"%s",
		"deadline":"%s"
	}`, createdUsername[0], createdPassword[0], taskname, deadline.Format(time.RFC3339)))
	request, err = http.NewRequest("POST", "/tasks", json_)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 500 {
		t.Error("task has been inserted although listId is not there")
	}
}

func TestGetTask(t *testing.T) {
	json_ := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, createdUsername[0], createdPassword[0]))
	request, err := http.NewRequest("GET", fmt.Sprintf("/tasks/%d", createdTaskId[0]), json_)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("task has not been fetched")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	var task data.Task
	err = json.Unmarshal(writer.Body.Bytes(), &task)
	if err != nil {
		return
	}
	if task.ListId != createdTaskListId[0] || task.Taskname != createdTaskname[0] ||
		task.Deadline != createdDeadline[0] || task.IsDone || task.IsImportant || task.Memo != "" {
		t.Error("task has not correctly been fetched")
	}
}

func TestUpdateTask(t *testing.T) {
	newTaskname := "new task name"
	newDeadline := time.Date(2022, time.April, 10, 0, 0, 0, 0, time.UTC)
	newMemo := "aiueo kkkkk"
	json_ := strings.NewReader(fmt.Sprintf(`{
		"username":"%s",
		"password":"%s",
		"listId":%d,
		"taskname":"%s",
		"deadline":"%s",
		"isDone":%s,
		"isImportant":%s, 
		"memo":"%s"
	}`, createdUsername[0], createdPassword[0], createdTaskListId[0], newTaskname, newDeadline.Format(time.RFC3339), "true", "true", newMemo))
	request, err := http.NewRequest("PUT", fmt.Sprintf("/tasks/%d", createdTaskId[1]), json_)
	if err != nil {
		t.Error(err)
	}
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("task has not been updated")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	// check if task is updated
	task, err := data.TaskRetrieve(createdTaskId[1])
	if err != nil {
		t.Error(err)
	}
	if task.ListId != createdTaskListId[0] || task.Taskname != newTaskname ||
		!task.Deadline.Equal(newDeadline) || !task.IsDone || !task.IsImportant || task.Memo != newMemo {
		t.Log(task)
		t.Error("task has not correctly been updated")
	}
}

func TestDeleteTask(t *testing.T) {
	t.Skip()
}
