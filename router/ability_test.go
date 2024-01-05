package router

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
	Init(r)

	payload := []byte(`phone_number=11111111111&password=123456&identity=警察&id_position=中国&username=3333&avatar=3.jpg`)
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
	Init(r)

	payload := []byte(`phone_number=19157692290&password=123456`)
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
	Init(r)

	payload := []byte(`{
		"author_id":222222,
		"title":"wzhnbwzhnbwzhnbwzhnbwzhnb article",
		"content":"wzhnbwzhnbwzhnbwzhnbwzhnb",
		"label":"one",
		"status":3
	}`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/createArticle", bytes.NewBuffer(payload))
	if err != nil {
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
	Init(r)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/getArticle", nil)
	if err != nil {
		t.Fatal(err)
	}
	params := make(url.Values)
	params.Add("id", "1")
	req.URL.RawQuery = params.Encode()

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.Code)
	}
}

func TestUpdateArticle(t *testing.T) {
	r := gin.Default()
	Init(r)

	payload := []byte(`{
    "id": 100,
    "author_id": 222222,
    "title": "wzhnbwzhnbwzhnbwzhnbwzhnb article",
    "content": "xjnbxjnbxjnbxjnb",
    "label": "one",
    "Status": 3,
    "ID": 0,
    "CreatedAt": "2024-01-03T21:18:46.61+08:00",
    "UpdatedAt": "2024-01-03T21:18:46.61+08:00",
    "DeletedAt": null
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
	Init(r)

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
	Init(r)

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

func TestCommentArticle(t *testing.T) {
	r := gin.Default()
	Init(r)

	payload := []byte(`article_id=3&commentator_id=7&content=wqeqwee`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/commentArticle", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestLikeArticle(t *testing.T) {
	r := gin.Default()
	Init(r)

	payload := []byte(`article_id=3&liker_id=7`)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/likeArticle", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestLikeGetByArticleID(t *testing.T) {
	r := gin.Default()
	Init(r)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/getArticleLike", nil)
	if err != nil {
		t.Fatal(err)
	}
	params := make(url.Values)
	params.Add("article_id", "3")
	req.URL.RawQuery = params.Encode()

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.Code)
	}
}

func TestCommentGetByArticleID(t *testing.T) {
	r := gin.Default()
	Init(r)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/getArticleComment", nil)
	if err != nil {
		t.Fatal(err)
	}
	params := make(url.Values)
	params.Add("article_id", "3")
	req.URL.RawQuery = params.Encode()

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.Code)
	}
}
