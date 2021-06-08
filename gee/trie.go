package gee

type node struct {
	pattern  []string
	value    string
	children map[string]*node
	isWild   bool
}

func makeNode(val string) *node {
	return &node{
		pattern:  nil,
		value:    val,
		children: make(map[string]*node),
		isWild:   len(val) != 0 && (val[0] == ':' || val[0] == '*'),
	}
}

func (n *node) matchChild(value string) *node {
	if nd, ok := n.children[value]; ok {
		return nd
	}
	return nil
}

func (n *node) matchChildWild(value string) *node {
	if nd, ok := n.children[value]; ok {
		return nd
	}
	for _, nd := range n.children {
		if nd.isWild {
			return nd
		}
	}
	return nil
}

type tree struct {
	root *node
}

func makeTree() *tree {
	return &tree{
		root: makeNode(""),
	}
}

func (tr *tree) insert(values []string) {
	curNode := tr.root
	for _, part := range values {
		target := curNode.matchChild(part)
		if target != nil {
			curNode = target
		} else {
			tmpNode := makeNode(part)
			curNode.children[part] = tmpNode
			curNode = tmpNode
		}
	}
	curNode.pattern = values
}

func (tr *tree) search(values []string) *node {
	curNode := tr.root
	for _, part := range values {
		target := curNode.matchChildWild(part)
		if target != nil {
			curNode = target
		} else {
			return nil
		}
	}
	return curNode
}
