package binarytree

import "fmt"

// ItemType - type of items
type ItemType int

// Tree struct
type Tree struct {
	root *nodeType
}

type nodeType struct {
	value ItemType
	left  *nodeType
	right *nodeType
}

// New - create new tree
func New() Tree {
	return Tree{
		root: nil,
	}
}

// Depth - return tree depth
func (tree *Tree) Depth() int {
	return treeDepth(tree.root)
}

func treeDepth(node *nodeType) int {
	if node != nil {
		leftDepth := treeDepth(node.left)
		rightDepth := treeDepth(node.right)

		switch leftDepth > rightDepth {
		case true:
			return leftDepth + 1
		case false:
			return rightDepth + 1
		default:
			// left and right depths is equal
			return leftDepth + 1
		}
	}

	return 0
}

// Add new node to tree
func (tree *Tree) Add(value ItemType) {
	if tree.root == nil {
		tree.root = &nodeType{
			value: value,
			left:  nil,
			right: nil,
		}
	} else {
		addNodeRecursively(tree.root, value)
	}
}

func addNodeRecursively(node *nodeType, value ItemType) *nodeType {
	if node == nil {
		return &nodeType{
			value: value,
			left:  nil,
			right: nil,
		}
	} else if value < node.value {
		node.left = addNodeRecursively(node.left, value)
	} else if value >= node.value {
		node.right = addNodeRecursively(node.right, value)
	}
	return node
}

// RemoveItem from tree
func (tree *Tree) RemoveItem(value ItemType) {
	removeItemRecursively(tree.root, value)
}

func removeItemRecursively(node *nodeType, value ItemType) *nodeType {
	return nil
}

// IsEmpty - reture bool value
func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

// Clear - delete all items from tree
func (tree *Tree) Clear() {
	// clear(tree.root)

	// GC will delete all unattainable nodes
	tree.root = nil
}

// Print - infix print
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
