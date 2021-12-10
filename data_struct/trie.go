package data_struct

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) insert(str string) {
	t.root.addNode(str, 0)
}

func (t *Trie) find(str string) bool {
	node := t.root.findNode(str, 0)
	return node != nil
}

type Node struct {
	Pattern   string  // origin string
	Part      uint8   // current char
	IsEndChar bool    // whether Part is the end of then current string
	Children  []*Node // next char
}

// matchChar match char
func (n *Node) matchCharNode(char uint8) *Node {
	for _, item := range n.Children {
		if item.Part == char {
			return item
		}
	}
	return nil
}

func (n *Node) matchCharNodes(char uint8) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Part == char {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *Node) addNode(pattern string, height int) {
	if len(pattern) == height {
		n.Pattern = pattern
		n.IsEndChar = true
		return
	}
	part := pattern[height]
	node := n.matchCharNode(part)
	if node == nil {
		node = &Node{Part: part}
		n.Children = append(n.Children, node)
	}
	node.addNode(pattern, height+1)
}

func (n *Node) findNode(pattern string, height int) *Node {
	if len(pattern) == height {
		if !n.IsEndChar {
			return nil
		}
		return n
	}

	part := pattern[height]
	nodes := n.matchCharNodes(part)
	for _, node := range nodes {
		findNode := node.findNode(pattern, height+1)
		if findNode != nil {
			return findNode
		}
	}
	return nil
}
