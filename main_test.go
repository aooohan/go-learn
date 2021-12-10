package main

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func Greet(writer io.Writer, name string) {
	//list.Element{}
}

func TestGreet(t *testing.T) {
	str := "你123123"
	for _, item := range str {
		fmt.Println(reflect.TypeOf(item))
	}
}
