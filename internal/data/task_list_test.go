package data

import (
	"testing"
)

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
