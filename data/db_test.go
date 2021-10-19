package data

import (
	"log"
	"testing"
	"time"
)

func TestJsonTime(t *testing.T) {
	var date string = "2011-01-02T10:02"
	dateByte := []byte(date)
	var jt JsonTime
	err := jt.UnmarshalJSON(dateByte)
	if err != nil {
		t.Error(err)
	}
	log.Println(jt)
}
func TestInsertDoneDelete(t *testing.T) {
	deadline := time.Date(2021, time.November, 1, 1, 0, 0, 0, time.UTC)
	task := Task{
		TaskName: "教科書を買う",
		Deadline: &JsonTime{deadline},
		IsDone:   false,
		DoneTime: TimeZero,
	}

	// insert
	err := task.Insert()
	if err != nil {
		t.Error(err)
	}
	id := task.Id

	var tasks []Task
	tasks, err = GetAllTask()
	if err != nil {
		t.Error(err)
	}

	for i := range tasks {
		log.Println("donetime", tasks[i].DoneTime)
		if tasks[i].Id == id {
			if tasks[i].TaskName != "教科書を買う" || tasks[i].IsDone != false {
				log.Println(tasks[i].TaskName)
				// log.Println(tasks[i].Deadline)
				log.Println(tasks[i].IsDone)
				t.Error("not correctly inserted")
			}
		}
	}

	// done
	err = DoneTask(id)
	if err != nil {
		t.Error(err)
	}

	tasks, err = GetAllTask()
	if err != nil {
		t.Error(err)
	}

	for i := range tasks {
		log.Println("donetime", tasks[i].DoneTime)
		if tasks[i].Id == id {
			if tasks[i].TaskName != "教科書を買う" || tasks[i].IsDone != true || tasks[i].DoneTime == TimeZero {
				log.Println(tasks[i].TaskName)
				// log.Println(tasks[i].Deadline)
				log.Println(tasks[i].IsDone)
				t.Error("not correctly done")
			}
		}
	}

	tasks, err = GetAllTask()
	if err != nil {
		t.Error(err)
	}

	for i := range tasks {
		if tasks[i].Id == id {
			if tasks[i].TaskName != "教科書を買う" || tasks[i].IsDone != true {
				t.Error("not correctly done")
			}
			// if tasks[i].DoneTime == time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC) {
			// 	t.Error("not correctly done")
			// }
		}
	}

	// delete
	err = DeleteTask(id)
	if err != nil {
		t.Error(err)
	}

	tasks, err = GetAllTask()
	if err != nil {
		t.Error(err)
	}

	for i := range tasks {
		if tasks[i].Id == id {
			t.Error("not correctly deleted")
		}
	}
}

func TestGetDoneTask(t *testing.T) {
	tasks, err := GetDoneTask()
	if err != nil {
		t.Error(err)
	}
	for i := range tasks {
		if tasks[i].IsDone != true {
			t.Error("get Doing Task")
		}
	}
}

func TestGetDoingTask(t *testing.T) {
	tasks, err := GetDoingTask()
	if err != nil {
		t.Error(err)
	}
	for i := range tasks {
		if tasks[i].IsDone != false {
			t.Error("get Done Task")
		}
	}
}
