package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello exec")
		wri, err := writer.Write([]byte("hello world"))
		if err != nil {
			return
		}
		fmt.Println(wri)
	})
	http.ListenAndServe("localhost:8080", nil)

}

func TestMap(t *testing.T) {
	m := make(map[string]string)
	k := "1231"
	s, err := m[k]
	if err {
		fmt.Println("输出：", s)
	}

}
