package route

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux http.ServeMux
var writer *httptest.ResponseRecorder

// REVIEW:よりよいテストの仕方,依存性の注入?
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = GetMux()
	writer = httptest.NewRecorder()
}

func TestCreateUser(t *testing.T) {
	// usernameを変えないとusernameの重複を許さないためテストが失敗する
	json := strings.NewReader(`{
		"username":"fagge9092",
		"password":"98h9ghiafe"
	}`)
	request, err := http.NewRequest("POST", "/users", json)
	if err != nil {
		t.Error(err)
	}
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Error("cannot create user")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}
}

func TestDeleteUser(t *testing.T) {
	json := strings.NewReader(`{
		"username":"fagfa92",
		"password":"98hfaiafe"
	}`)
	request, _ := http.NewRequest("POST", "/users", json)
	writer = httptest.NewRecorder() // REVIEW:NewRecorderは生成しなくていいはず、bufferを消すとかで使いまわせそう
	mux.ServeHTTP(writer, request)

	json = strings.NewReader(`{
		"username":"fagfa92",
		"password":"98hfaiafe"
	}`)
	request, _ = http.NewRequest("DELETE", "/users", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Error("cannot delete user")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}
}

func TestLogin(t *testing.T) {
	json := strings.NewReader(`{
		"username":"faoj",
		"password":"9faade"
	}`)
	request, _ := http.NewRequest("POST", "/users", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)

	// ログインできること
	json = strings.NewReader(`{
		"username":"faoj",
		"password":"9faade"
	}`)
	request, _ = http.NewRequest("POST", "/users/login", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 201 {
		t.Error("cannot login user")
		errMsg := writer.Body.String()
		t.Error(errMsg)
	}

	// passwordが違うとログインできないこと
	json = strings.NewReader(`{
		"username":"faoj",
		"password":"ohgia"
	}`)
	request, _ = http.NewRequest("POST", "/users/login", json)
	writer = httptest.NewRecorder()
	mux.ServeHTTP(writer, request)
	if writer.Code != 500 {
		t.Fatal("can login although password is not correct")
	}
}
