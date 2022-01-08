package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshall(t *testing.T) {
	taskList := TaskList{
		ListId:   1443,
		Username: "aiueo",
		Icon:     "sub",
		Listname: "aieogb",
	}
	jsona, err := json.Marshal(taskList)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(jsona), `{"listId":1443,"username":"aiueo","icon":"sub","listname":"aieogb"}`)

	taskList = TaskList{
		ListId:   1443,
		Icon:     "sub",
		Listname: "aieogb",
	}
	jsonb, err := json.Marshal(taskList)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(jsonb), `{"listId":1443,"icon":"sub","listname":"aieogb"}`)
}

// REVIEW: よりよいdbまわりのテストの仕方,依存をなくす
func TestTaskListCRUD(t *testing.T) {
	username := "hedwig100"
	icon := "add"
	listname := "mid-term test"

	UserCreate(username, "password") // ここではこのエラーは無視(user_testでテストする)

	// create
	taskList, err := TaskListCreate(username, icon, listname)
	if err != nil {
		t.Error(err)
	}
	if taskList.Username != username || taskList.Icon != icon || taskList.Listname != listname {
		t.Fatalf("taskList: %v", taskList)
	}

	// retrive
	taskListR, err := TaskListRetrieve(taskList.ListId)
	if err != nil {
		t.Error(err)
	}
	if taskListR.ListId != taskList.ListId || taskListR.Username != taskList.Username ||
		taskListR.Icon != taskList.Icon || taskListR.Listname != taskList.Listname {
		t.Fatalf("taskList: %v", taskList)
	}

	// update
	taskList.Listname = "term test"
	err = TaskListUpdate(taskList)
	if err != nil {
		t.Error(err)
	}
	taskListR, err = TaskListRetrieve(taskList.ListId)
	if err != nil {
		t.Error(err)
	}
	if taskListR.ListId != taskList.ListId || taskListR.Username != taskList.Username ||
		taskListR.Icon != taskList.Icon || taskListR.Listname != taskList.Listname {
		t.Fatalf("taskList: %v", taskList)
	}

	// delete
	err = TaskListDelete(taskList.ListId)
	if err != nil {
		t.Error(err)
	}
	_, err = TaskListRetrieve(taskList.ListId)
	if err == nil {
		t.Error("cannot delete task_list")
	}
}

func TestTaskListAndTasks(t *testing.T) {
	tasklist, tasks, err := TaskListAndTasks(createdTaskListId[0])
	if err != nil {
		t.Error(err)
	}

	if tasklist.ListId != createdTaskListId[0] || tasklist.Username != createdUsername[0] ||
		tasklist.Icon != "add" || tasklist.Listname != createdTaskListname[0] {
		t.Error("tasklist is not correctly retrieved")
	}

	if len(tasks) != len(createdTaskId) {
		t.Error("the number of tasks is not correct")
	}

	for index, task := range tasks {
		if task.TaskId != createdTaskId[index] || task.Taskname != createdTaskname[index] ||
			task.IsDone != false || task.IsImportant != false || task.Memo != "" {
			t.Error("task is not correctly retrived")
		}
	}
}
