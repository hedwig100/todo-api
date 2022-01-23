package data

import (
	"testing"
	"time"
)

// REVIEW: よりよいdbまわりのテストの仕方,依存をなくす
func TestTaskCRUD(t *testing.T) {
	username := "hedwig100"
	taskname := "meet gopher "
	deadline := time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)
	listId := createdTaskListId[1]

	// create
	task, err := TaskCreate(username, listId, taskname, deadline)
	if err != nil {
		t.Error(err)
	}
	if task.Username != username || task.ListId != listId || task.Taskname != taskname ||
		task.Deadline != deadline || task.IsDone == true || task.IsImportant == true {
		t.Fatalf("task: %v", task)
	}

	// retrive
	taskR, err := TaskRetrieve(task.TaskId)
	if err != nil {
		t.Error(err)
	}
	if taskR.Username != task.Username || taskR.ListId != task.ListId || taskR.Taskname != task.Taskname ||
		!task.Deadline.Equal(taskR.Deadline) || taskR.IsDone != task.IsDone || taskR.IsImportant != task.IsImportant {
		t.Logf("orginal task: %v", task)
		t.Fatalf("retrieved task: %v", taskR)
	}

	// update
	task.IsDone = true
	task.Taskname = "study math"
	err = TaskUpdate(task)
	if err != nil {
		t.Error(err)
	}
	taskR, err = TaskRetrieve(task.TaskId)
	if err != nil {
		t.Error(err)
	}
	if taskR.Username != task.Username || taskR.ListId != task.ListId || taskR.Taskname != task.Taskname ||
		!task.Deadline.Equal(taskR.Deadline) || taskR.IsDone != task.IsDone || taskR.IsImportant != task.IsImportant {
		t.Logf("orginal task: %v", task)
		t.Fatalf("retrieved task: %v", taskR)
	}

	// delete
	err = TaskDelete(task.TaskId)
	if err != nil {
		t.Error(err)
	}
	_, err = TaskRetrieve(task.TaskId)
	if err == nil {
		t.Error("cannot delete task")
	}
}
