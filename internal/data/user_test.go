package data

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"testing"
)

const letters = "abcedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const letLen = 62

func makeRandomUsername(len int) string {
	usernameArray := make([]string, len)
	for i := 0; i < len; i++ {
		usernameArray[i] = string(letters[rand.Intn(letLen)])
	}
	return strings.Join(usernameArray, "")
}

// REVIEW: よりよいdbまわりのテストの仕方
func TestUserCRUD(t *testing.T) {
	username := makeRandomUsername(10)
	password := "aiueoABCED"

	// create
	user, err := UserCreate(username, password)
	if err != nil {
		t.Error(err)
	}

	flag := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if flag != nil || user.Username != username {
		t.Fatalf("user: %v", user)
	}

	// retrive
	user_db, err := UserRetrieve(username)
	if err != nil {
		t.Error(err)
	}

	flag = bcrypt.CompareHashAndPassword(user_db.Password, []byte(password))
	if flag != nil || user_db.Uuid != user.Uuid || user_db.Username != user.Username {
		t.Logf("original_user: %v", user)
		t.Fatalf("retrieved_user: %v", user_db)
	}

	// delete
	err = UserDelete(username)
	if err != nil {
		t.Error(err)
	}
	_, err = UserRetrieve(username)
	if err == nil {
		t.Error("cannot delete user")
	}
}
