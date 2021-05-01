package main

import (
	"fmt"
)

const MAXINT = int(^uint(0) >> 1)

var (
	graph [][]*mapCell
	mapSize int
)

type mapCell struct {
	Cost, Distance int
	X, Y           int
}

type PriorityQueue struct {
	Heap []*mapCell
	Size int
}

func (pq *PriorityQueue) Less(i, j int) bool {
	h := pq.Heap
	return h[i].Distance < h[j].Distance
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Heap[i], pq.Heap[j] = pq.Heap[j], pq.Heap[i]
}

func (pq *PriorityQueue) Push(v *mapCell) {
	pq.Heap = append(pq.Heap, v)
	pq.Size++
	pq.SiftUp(pq.Size-1)
}

func (pq *PriorityQueue) Pop() *mapCell {
	res := pq.Heap[0]
	pq.Size--
	pq.Heap = pq.Heap[1:]
	return res
}

func (pq *PriorityQueue) SiftUp(i int) {
	for ; i != 0 && pq.Less(i, (i-1)/2); i = (i - 1) / 2 {
		pq.Swap(i, (i-1)/2)
	}
}

func (pq *PriorityQueue) SiftDown(i int) {
	for i < pq.Size/2 {
		mi := i*2 + 1
		if mi + 1 < pq.Size && pq.Less(mi + 1, mi) {
			mi++
		}
		if pq.Less(i, mi) {
			return
		}
		pq.Swap(i, mi)
		i = mi
	}
}

func dijkstra() int {
	var nextCell, current *mapCell
	pq := &PriorityQueue{make([]*mapCell, 0, mapSize), 0}
	graph[0][0].Distance = 0
	pq.Push(graph[0][0])
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for pq.Size != 0 {
		current = pq.Pop()
		for i := 0; i < 4; i++ {
			if nextX, nextY := current.X+dx[i], current.Y+dy[i]; nextX < mapSize && nextX >= 0 && nextY < mapSize && nextY >= 0 {
				nextCell = graph[nextX][nextY]
				if newDistance := nextCell.Cost + current.Distance; newDistance < nextCell.Distance {
					nextCell.Distance = newDistance
					pq.Push(nextCell)
				}
			}
		}
	}

	return graph[mapSize-1][mapSize-1].Distance + graph[0][0].Cost
}

func main() {
	var w int
	fmt.Scanf("%d", &mapSize)

	graph = make([][]*mapCell, mapSize)
	for i := 0; i < mapSize; i++ {
		graph[i] = make([]*mapCell, mapSize)
	}
	for i := 0; i < mapSize; i++ {
		for j := 0; j < mapSize; j++ {
			fmt.Scanf("%d", &w)
			graph[i][j] = &mapCell{w, MAXINT, i, j}
		}
	}

	fmt.Printf("%d", dijkstra())
}
