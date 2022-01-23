package data

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uuid     int
	Username string
	Password []byte
}

func UserCreate(username string, password string) (user User, err error) {
	// パスワードをハッシュ化する
	hashedPassword := []byte(password)
	hashedPassword, err = bcrypt.GenerateFromPassword(hashedPassword, 10)

	if err != nil {
		return
	}

	user = User{
		Username: username,
		Password: hashedPassword,
	}

	stmt, err := Db.Prepare("INSERT INTO users (username,password) VALUES ($1,$2) RETURNING uuid")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Username, user.Password).Scan(&user.Uuid)
	return
}

func UserRetrieve(username string) (user User, err error) {
	stmt, err := Db.Prepare("SELECT uuid,password FROM users WHERE username = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	user = User{Username: username}
	err = stmt.QueryRow(username).Scan(&user.Uuid, &user.Password)
	return
}

func UserDelete(username string) (err error) {
	stmt, err := Db.Prepare("DELETE FROM users WHERE username = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	return
}

func Login(username string, password string) (user User, success bool, err error) {
	user, err = UserRetrieve(username)
	if err != nil {
		return
	}

	if bcrypt.CompareHashAndPassword(user.Password, []byte(password)) != nil {
		success = false
	} else {
		success = true
	}
	return
}

func UsersTaskLists(username string) (taskLists []TaskList, err error) {
	stmt, err := Db.Prepare("SELECT list_id,icon,listname FROM task_lists WHERE username = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return
	}

	for rows.Next() {
		tasklist := TaskList{Username: username}
		err = rows.Scan(&tasklist.ListId, &tasklist.Icon, &tasklist.Listname)
		if err != nil {
			return
		}
		taskLists = append(taskLists, tasklist)
	}

	err = rows.Err()
	return
}
