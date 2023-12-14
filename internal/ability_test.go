package internal

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRegister(t *testing.T) {
	payload := []byte(`username=wzh&password=123456`)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

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

	recorder := httptest.NewRecorder()
	UserLogin(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestArticleCreate(t *testing.T) {
	payload := []byte(`{"userid"="1111","text"="wzhnbwzhnbwzhnb"}`)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	ArticleCreate(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

//func TestArticleGet(t *testing.T) {
//	req, err := http.NewRequest(http.MethodGet, "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	params := make(url.Values)
//	params.Add("id", "1111")
//	req.URL.RawQuery = params.Encode()
//	recorder := httptest.NewRecorder()
//
//	ArticleGet(recorder, req)
//
//	if recorder.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
//	}
//}
//
//func TestArticleUpdate(t *testing.T) {
//	req, err := http.NewRequest(http.MethodDelete, "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	params := make(url.Values)
//	params.Add("id", "1111")
//	req.URL.RawQuery = params.Encode()
//	recorder := httptest.NewRecorder()
//
//	ArticleGet(recorder, req)
//
//	if recorder.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
//	}
//}

func TestArticleDelete(t *testing.T) {

}
