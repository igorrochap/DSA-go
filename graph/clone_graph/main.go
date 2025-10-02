package main

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

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }

    var q Dequeue[*Node]
    q.PushFront(node)
    clones := map[int]*Node{}
    clones[node.Val] = &Node{node.Val, []*Node{}}

    for q.Len() > 0 {
        curr, _ := q.PopFront()
        currClone := clones[curr.Val]

        for _, n := range curr.Neighbors {
            if _, ok := clones[n.Val]; !ok {
                clones[n.Val] = &Node{n.Val, []*Node{}}
                q.PushBack(n)
            }
            currClone.Neighbors = append(currClone.Neighbors, clones[n.Val])
        }
    }
    return clones[node.Val]
}