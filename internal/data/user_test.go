package data

import (
	"log"
	"sort"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// テストデータ
var createdUsername = []string{"hedwig100", "pokemon", "mac"}
var createdPassword = []string{"iajgo3o", ")8hgiau", "uhaig1928"}
var createdTaskListname = []string{"cooking for chistmas", "for presentation", "mid-term test"}
var createdTaskListId []int

func TestMain(m *testing.M) {
	err := setUp()
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func setUp() (err error) {
	// dbの初期化
	_, err = Db.Exec("DELETE FROM users")
	if err != nil {
		return
	}

	// テストデータ挿入
	for index, username := range createdUsername {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createdPassword[index]), 10)
		if err != nil {
			return err
		}
		Db.Exec("INSERT INTO users (username,password) VALUES ($1,$2)", username, hashedPassword)
	}
	createdTaskListId = make([]int, 3)
	for index, listname := range createdTaskListname {
		var listId int
		err = Db.QueryRow(`INSERT INTO task_lists (username,icon,listname) VALUES ($1,'add',$2) RETURNING list_id`,
			createdUsername[0], listname).Scan(&listId)
		if err != nil {
			return
		}
		createdTaskListId[index] = listId
	}
	return
}

// REVIEW: よりよいdbまわりのテストの仕方

func TestUserRetrieve(t *testing.T) {
	user, err := UserRetrieve(createdUsername[0])
	if err != nil {
		t.Error(err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(createdPassword[0]))
	if err != nil || user.Username != createdUsername[0] {
		t.Error(err)
		t.Error("username is not correct")
	}
}

func TestUserCreate(t *testing.T) {
	_, err := UserCreate("aiueoe", "giohau")
	if err != nil {
		t.Error(err)
	}
	user, err := UserRetrieve("aiueoe")
	if err != nil {
		t.Error(err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte("giohau"))
	if err != nil || user.Username != "aiueoe" {
		t.Error(err)
		t.Error("username is not correct")
	}
}

func TestUserDelete(t *testing.T) {
	err := UserDelete(createdUsername[1])
	if err != nil {
		t.Error(err)
	}

	_, err = UserRetrieve(createdUsername[1])
	if err == nil {
		t.Error("cannot delete user")
	}
}

func TestUsersTaskList(t *testing.T) {
	taskLists, err := UsersTaskLists(createdUsername[0])
	if err != nil {
		t.Error(err)
	}

	sort.Slice(taskLists, func(i, j int) bool { return taskLists[i].Listname < taskLists[j].Listname })
	for index, tasklist := range taskLists {
		if tasklist.ListId != createdTaskListId[index] ||
			tasklist.Listname != createdTaskListname[index] ||
			tasklist.Username != createdUsername[0] {
			t.Error("cannot get user's tasklists")
		}
	}
}
