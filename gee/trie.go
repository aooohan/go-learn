package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // route path
	part     string  // a part of route path
	children []*node // children nodes
	isWild   bool    // is ":xxx" or "*"
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// matchChild find first match node,use for insert
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren find all nodes which match part,use for search
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert :insert a node for trie tree
func (n *node) insert(pattern string, parts []string, height int) {
	if height == len(parts) {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 到底， 或者当前这个part 以 * 开头
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		// only the pattern field of the leaf node has value
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
