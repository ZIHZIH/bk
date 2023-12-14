package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	payload := []byte(`username=wzh&password=123456`)
	request, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/user/login", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(request)
	result, _ := io.ReadAll(resp.Body)
	fmt.Println(string(result))
}

//func RegLoginHandler(f func(writer http.ResponseWriter, request *http.Request) error) func(writer http.ResponseWriter, request *http.Request) {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Println("login before")
//		err := f(writer, request)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("login after")
//	}
//}
