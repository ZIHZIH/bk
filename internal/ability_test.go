package internal

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserRegister(t *testing.T) {
	payload := []byte(`username=wzh&password=123456`)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = req.Body.Close() }()

	recorder := httptest.NewRecorder()

	UserRegister(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestUserLogin(t *testing.T) {
	payload := []byte(`username=wzh&password=123456`)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = req.Body.Close() }()

	recorder := httptest.NewRecorder()
	UserLogin(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestArticleCreate(t *testing.T) {
	payload := []byte(`{"userid"="1111","text"="wzhnbwzhnbwzhnb"}`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/article", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = req.Body.Close() }()

	recorder := httptest.NewRecorder()
	ArticleCreate(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestArticleGet(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/article", nil)
	if err != nil {
		t.Fatal(err)
	}

	params := make(url.Values)
	params.Add("id", "1")
	req.URL.RawQuery = params.Encode()
	recorder := httptest.NewRecorder()

	ArticleGet(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestArticleUpdate(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	b := []byte(`{"id":"3","Article":{"userid":"33333","text":"cccccccc"}}`)
	req, err := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

	params := make(url.Values)
	params.Add("id", "1111")
	req.URL.RawQuery = params.Encode()
	recorder := httptest.NewRecorder()

	ArticleGet(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestArticleDelete(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	req, err := http.NewRequest(http.MethodDelete, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = req.Body.Close() }()
	params := make(url.Values)
	params.Add("id", "1")
	req.URL.RawQuery = params.Encode()
	recorder := httptest.NewRecorder()

	ArticleDelete(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
