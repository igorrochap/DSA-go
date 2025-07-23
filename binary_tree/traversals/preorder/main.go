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
	tree.insertRecursive(value, tree.Root)
}

func (tree *BinaryTree) insertRecursive(value int, node *Node) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			tree.insertRecursive(value, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			tree.insertRecursive(value, node.Right)
		}
	}
}

func (tree *BinaryTree) Search(value int) bool {
	return tree.searchRecursive(tree.Root, value)
}

func (tree *BinaryTree) searchRecursive(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	} else if value < node.Value {
		return tree.searchRecursive(node.Left, value)
	} else {
		return tree.searchRecursive(node.Right, value)
	}
}

func (tree *BinaryTree) PreorderTraversal() []int {
	var result []int
	tree.preorderRecursive(tree.Root, &result)
	return result
}

func (tree *BinaryTree) preorderRecursive(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		tree.preorderRecursive(node.Left, result)
		tree.preorderRecursive(node.Right, result)
	}
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(7)

	fmt.Println(tree.PreorderTraversal())
}
