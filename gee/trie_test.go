package gee

import (
	"fmt"
	"strings"
	"testing"
)

func TestInsert(t *testing.T) {
	node := &node{children: make([]*node, 0)}
	p1 := "a/e/c/d"
	p2 := "a/b/*filename"
	node.insert(p1, strings.Split(p1, "/"), 0)
	node.insert(p2, strings.Split(p2, "/"), 0)
	search := node.search(strings.Split("a/b/c/d", "/"), 0)
	fmt.Println(search)
	fmt.Println(node)
}
