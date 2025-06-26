package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value}
		return
	}
	tree.insertRecursively(value, tree.Root)
}

func (tree *BinaryTree) insertRecursively(value int, node *Node) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			tree.insertRecursively(value, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			tree.insertRecursively(value, node.Right)
		}
	}
}

func (tree *BinaryTree) Search(value int) bool {
	return tree.searchRecursively(tree.Root, value)
}

func (tree *BinaryTree) searchRecursively(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	} else if value < node.Value {
		return tree.searchRecursively(node.Left, value)
	} else {
		return tree.searchRecursively(node.Right, value)
	}
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(9)

	fmt.Printf("Search 4: %v\n", tree.Search(4))
	fmt.Printf("Search 6: %v\n", tree.Search(6))
}
