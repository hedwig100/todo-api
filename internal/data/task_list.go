package data

type TaskList struct {
	ListId   int
	Username string
	Icon     string
	Listname string
}

func TaskListCreate(username string, icon string, listname string) (taskList TaskList, err error) {
	taskList = TaskList{
		Username: username,
		Icon:     icon,
		Listname: listname,
	}

	stmt, err := Db.Prepare("INSERT INTO task_lists (username,icon,listname) VALUES ($1,$2,$3) RETURNING list_id")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(taskList.Username, taskList.Icon, taskList.Listname).Scan(&taskList.ListId)
	return
}

func TaskListRetrieve(listId int) (taskList TaskList, err error) {
	stmt, err := Db.Prepare("SELECT username,icon,listname FROM task_lists WHERE list_id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	taskList = TaskList{ListId: listId}
	err = stmt.QueryRow(taskList.ListId).Scan(&taskList.Username, &taskList.Icon, &taskList.Listname)
	return
}

func TaskListUpdate(taskList TaskList) (err error) {
	stmt, err := Db.Prepare("UPDATE task_lists SET username = $1,icon = $2,listname = $3 WHERE list_id = $4")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskList.Username, taskList.Icon, taskList.Listname, taskList.ListId)
	return
}

func TaskListDelete(listId int) (err error) {
	stmt, err := Db.Prepare("DELETE FROM task_lists WHERE list_id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(listId)
	return
}
