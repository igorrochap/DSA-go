package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	Value    string
	Distance int
	Index    int
}

type HeapQueue []*Item

func (hp HeapQueue) Len() int {
	return len(hp)
}

func (hp HeapQueue) Less(i, j int) bool {
	return hp[i].Distance < hp[j].Distance
}

func (hp HeapQueue) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
	hp[i].Index = i
	hp[j].Index = j
}

func (hp *HeapQueue) Push(item interface{}) {
	n := len(*hp)
	newItem := item.(*Item)
	newItem.Index = n
	*hp = append(*hp, newItem)
}

func (hp *HeapQueue) Pop() interface{} {
	old := *hp
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*hp = old[0 : n-1]
	return item
}

type Graph map[string]map[string]int

func dijkstra(graph Graph, start string) map[string]int {
	minHeap := make(HeapQueue, 0)
	heap.Init(&minHeap)
	minHeap.Push(&Item{Value: start, Distance: 0})

	distances := make(map[string]int)
	for node := range graph {
		distances[node] = int(math.Inf(1))
	}
	distances[start] = 0

	for minHeap.Len() > 0 {
		currNode := minHeap.Pop().(*Item)
		if currNode.Distance > distances[currNode.Value] {
			continue
		}
		for neighbor, weight := range graph[currNode.Value] {
			distance := currNode.Distance + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				minHeap.Push(&Item{Value: neighbor, Distance: distance})
			}
		}
	}

	return distances
}

func main() {
	graph := Graph{
		"A": {"B": 1, "C": 4},
		"B": {"A": 1, "C": 2, "D": 5},
		"C": {"A": 4, "B": 2, "D": 1},
		"D": {"B": 5, "C": 1},
	}
	shortestPaths := dijkstra(graph, "A")
	fmt.Println(shortestPaths)
}
