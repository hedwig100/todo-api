package data

import (
	"time"
)

type Task struct {
	TaskId      int
	Username    string
	ListId      int
	Taskname    string
	Deadline    time.Time
	IsDone      bool
	IsImportant bool
	Memo        string
}

func TaskCreate(username string, listId int, taskname string, deadline time.Time) (task Task, err error) {
	task = Task{
		Username:    username,
		ListId:      listId,
		Taskname:    taskname,
		Deadline:    deadline,
		IsDone:      false,
		IsImportant: false,
	}

	stmt, err := Db.Prepare("INSERT INTO tasks (username,list_id,taskname,deadline) VALUES ($1,$2,$3,$4) RETURNING task_id")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(task.Username, task.ListId, task.Taskname, task.Deadline).Scan(&task.TaskId)
	return
}

func TaskRetrieve(taskId int) (task Task, err error) {
	stmt, err := Db.Prepare("SELECT username,list_id,taskname,deadline,is_done,is_important,memo FROM tasks WHERE task_id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	task = Task{TaskId: taskId}
	err = stmt.QueryRow(taskId).Scan(&task.Username, &task.ListId, &task.Taskname, &task.Deadline, &task.IsDone, &task.IsImportant, &task.Memo)
	return
}

func TaskUpdate(task Task) (err error) {
	stmt, err := Db.Prepare("UPDATE tasks SET username = $1,list_id = $2,taskname = $3,deadline = $4,is_done = $5,is_important = $6,memo = $7 WHERE task_id = $8")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Username, task.ListId, task.Taskname, task.Deadline, task.IsDone, task.IsImportant, task.Memo, task.TaskId)
	return
}

func TaskDelete(taskId int) (err error) {
	stmt, err := Db.Prepare("DELETE FROM tasks WHERE task_id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskId)
	return
}
