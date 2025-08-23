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
	tree.insertRecursive(tree.Root, value)
}

func (tree *BinaryTree) insertRecursive(node *Node, value int) {
	if node.Value > value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			tree.insertRecursive(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			tree.insertRecursive(node.Right, value)
		}
	}
}

func (tree *BinaryTree) DFS(value int) bool {
	return tree.DFSRecursive(tree.Root, value)
}

func (tree *BinaryTree) DFSRecursive(node *Node, value int) bool {
	if node == nil {
		return false
	}
	fmt.Println(node.Value)
	if node.Value == value {
		return true
	}
	if tree.DFSRecursive(node.Left, value) {
		return true
	}
	if tree.DFSRecursive(node.Right, value) {
		return true
	}
	return false
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(20)

	fmt.Println(tree.DFS(20))
}
