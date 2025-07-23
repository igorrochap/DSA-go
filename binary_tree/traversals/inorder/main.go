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
	if value < node.Value {
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

func (tree *BinaryTree) InorderTraversal() []int {
	var result []int
	tree.inorderRecursive(tree.Root, &result)
	return result
}

func (tree *BinaryTree) inorderRecursive(node *Node, result *[]int) {
	if node != nil {
		tree.inorderRecursive(node.Left, result)
		*result = append(*result, node.Value) // adds to the result only after all nodes in the left were visited
		tree.inorderRecursive(node.Right, result)
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

	fmt.Println(tree.InorderTraversal())
}
