package internal

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserRegister(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	payload := []byte(`phone_number=wzh&password=123456&identity=警察&id_position=中国`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/user/register", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestUserLogin(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	payload := []byte(`phone_number=wzh&password=123456`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/user/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
	}
}

func TestCreateArticle(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	payload := []byte(`{
	   "user_id":222222,
	   "title":"wzhnbwzhnbwzhnbwzhnbwzhnb article",
	   "content":"wzhnbwzhnbwzhnbwzhnbwzhnb",
	   "label":"wzhnbwzhnbwzhnbwzhnbwzhnb",
	   "status":3
	}`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/createArticle", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("qweqweqweeqweqweqweqweqweqwewqeqweqwewq")
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestGetArticle(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/getArticle", nil)
	if err != nil {
		t.Fatal(err)
	}
	params := make(url.Values)
	params.Add("id", "10")
	req.URL.RawQuery = params.Encode()

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.Code)
	}
}

func TestUpdateArticle(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	payload := []byte(`{
    "id": 2,
    "user_id": 332323,
    "title": "wzhnbwzhnbwzhnbwzhnbwzhnb article",
    "content": "wzhnbwzhnbwzhnbwzhnbwzhnb",
    "label": "wzhnbwzhnbwzhnbwzhnbwzhnb",
    "status": 99,
    "create_time": "2023-12-20T13:50:41.517585+08:00"
}`)
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/updateArticle", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("qweqweqweeqweqweqweqweqweqwewqeqweqwewq")
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestDeleteArticle(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	req, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/deleteArticle", nil)
	if err != nil {
		t.Fatal(err)
	}
	params := make(url.Values)
	params.Add("id", "1")
	req.URL.RawQuery = params.Encode()

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestListArticle(t *testing.T) {
	r := gin.Default()
	InitGinRouter(r)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/listArticle", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}
