package data_struct

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	trie := NewTrie()
	trie.insert("abcde")
	trie.insert("abc")
	trie.insert("bcd")
	find := trie.find("abc")
	if !find {
		t.Error("error")
	}

	f2 := trie.find("bcd")
	if !f2 {
		t.Error("error")
	}
	fmt.Println(find)
	fmt.Println(f2)
}

func TestInsertHanzi(t *testing.T) {
	trie := NewTrie()
	trie.insert("你好")
	trie.insert("你好啊世界")
	trie.insert("明天你好")
	exist := trie.find("你好")
	if !exist {
		t.Error("error")
	}
}
