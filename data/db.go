package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Task struct {
	Id       int      `json:"id"`
	TaskName string   `json:"taskname"`
	Deadline JsonTime `json:"deadline"`
	IsDone   bool     `json:"isdone"`
	DoneTime JsonTime `json:"donetime"`
}

var Db *sql.DB
var TimeZero JsonTime

// var LOC *time.Location

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=todo dbname=todo password=todo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	TimeZero = JsonTime{time.Time{}}
	// LOC, err = time.LoadLocation("Asia/Tokyo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// insert task to db
func (task *Task) Insert() (err error) {
	statement := "INSERT INTO task (taskname, deadline, isdone,donetime) VALUES ($1,$2,$3,$4) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(task.TaskName, task.Deadline.JsonTime, task.IsDone, task.DoneTime.JsonTime).Scan(&task.Id)
	if err != nil {
		return
	}
	return
}

func DoneTask(id int) (err error) {
	nowTime := time.Now()
	_, err = Db.Exec("UPDATE task SET isdone = TRUE,donetime = $1 WHERE id = $2", nowTime, id)
	return
}

// delete task
func DeleteTask(id int) (err error) {
	_, err = Db.Exec("DELETE FROM task WHERE id = $1", id)
	return
}

// get all task from db
func GetAllTask() (alltask []Task, err error) {
	rows, err := Db.Query("SELECT * FROM task")
	if err != nil {
		return
	}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.TaskName, &task.Deadline.JsonTime, &task.IsDone, &task.DoneTime.JsonTime)
		if err != nil {
			return
		}
		alltask = append(alltask, task)
	}
	defer rows.Close()
	return
}

// get done task
func GetDoneTask() (donetask []Task, err error) {
	rows, err := Db.Query("SELECT * FROM task WHERE isdone = TRUE")
	if err != nil {
		return
	}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.TaskName, &task.Deadline.JsonTime, &task.IsDone, &task.DoneTime.JsonTime)
		if err != nil {
			return
		}
		donetask = append(donetask, task)
	}
	defer rows.Close()
	return
}

// get doing task
func GetDoingTask() (doingtask []Task, err error) {
	rows, err := Db.Query("SELECT * FROM task WHERE isdone = FALSE")
	if err != nil {
		return
	}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.TaskName, &task.Deadline.JsonTime, &task.IsDone, &task.DoneTime.JsonTime)
		if err != nil {
			return
		}
		doingtask = append(doingtask, task)
	}
	defer rows.Close()
	return
}
