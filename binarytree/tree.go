package binarytree

import "fmt"

// Tree ...
type Tree struct {
	root *nodeType
}

type nodeType struct {
	value int
	left  *nodeType
	right *nodeType
}

// New ...
func New() Tree {
	return Tree{
		root: nil,
	}
}

// Add ...
func (tree *Tree) Add(value int) {
	if tree.root == nil {
		tree.root = &nodeType{
			value: value,
			left:  nil,
			right: nil,
		}
	} else {
		addNode(tree.root, value)
	}
}

func addNode(node *nodeType, value int) *nodeType {
	if node == nil {
		return &nodeType{
			value: value,
			left:  nil,
			right: nil,
		}
	} else if value < node.value {
		node.left = addNode(node.left, value)
	} else if value >= node.value {
		node.right = addNode(node.right, value)
	}

	return node
}

// Clear ...
func (tree *Tree) Clear() {
	// clear(tree.root)

	// gc will delete all unattainable nodes
	tree.root = nil
}

// func clear(node *nodeType) {
// 	if node != nil {
// 		clear(node.left)
// 		clear(node.right)
// 		node = nil
// 	}
// }

// Print ...
func (tree *Tree) Print() {
	infixPrint(tree.root)
}

// PrefixPrint ...
func (tree *Tree) PrefixPrint() {
	prefixPrint(tree.root)
}

func prefixPrint(node *nodeType) {
	if node != nil {
		fmt.Printf("%v\n", node.value)
		prefixPrint(node.left)
		prefixPrint(node.right)
	}
}

// InfixPrint ...
func (tree *Tree) InfixPrint() {
	infixPrint(tree.root)
}

func infixPrint(node *nodeType) {
	if node != nil {
		infixPrint(node.left)
		fmt.Printf("%v\n", node.value)
		infixPrint(node.right)
	}
}

// PostfixPrint ...
func (tree *Tree) PostfixPrint() {
	postfixPrint(tree.root)
}

func postfixPrint(node *nodeType) {
	if node != nil {
		postfixPrint(node.left)
		postfixPrint(node.right)
		fmt.Printf("%v\n", node.value)
	}
}
