package main

import "fmt"

type Dequeue[Item any] struct {
	Items []Item
}

func (queue *Dequeue[Item]) PushBack(value Item) {
	queue.Items = append(queue.Items, value)
}

func (queue *Dequeue[Item]) PushFront(value Item) {
	queue.Items = append([]Item{value}, queue.Items...)
}

func (queue *Dequeue[Item]) Len() int {
	return len(queue.Items)
}

func (queue *Dequeue[Item]) PopBack() (Item, bool) {
	if queue.Len() == 0 {
		var notFound Item
		return notFound, false
	}
	item := queue.Items[len(queue.Items)-1]
	queue.Items = queue.Items[:len(queue.Items)-1]
	return item, true
}

func (queue *Dequeue[Item]) PopFront() (Item, bool) {
	if queue.Len() == 0 {
		var notFound Item
		return notFound, false
	}
	item := queue.Items[0]
	queue.Items = queue.Items[1:]
	return item, true
}

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

func (tree *BinaryTree) BFS(value int) bool {
	if tree.Root == nil {
		return false
	}
	queue := &Dequeue[*Node]{}
	queue.PushBack(tree.Root)
	for queue.Len() > 0 {
		node, _ := queue.PopFront()
		fmt.Println(node.Value)
		if node.Value == value {
			return true
		}
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
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

	fmt.Println(tree.BFS(10))
}
