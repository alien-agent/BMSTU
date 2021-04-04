package main

import (
	"fmt"
	"math"
	"sort"
)

var (
	parentIndex   []int
	depth         []int
	trails        []Trail
	verticesCount int
)

type Vertex struct {
	x, y int
}

type Trail struct {
	FromIndex, ToIndex int
	Length             int
}

func Union(v1Index, v2Index int) {
	rootIndex1, rootIndex2 := FindParentIndex(v1Index), FindParentIndex(v2Index)
	if depth[rootIndex1] < depth[rootIndex2] {
		parentIndex[rootIndex1] = rootIndex2
	} else {
		parentIndex[rootIndex2] = rootIndex1
		if (depth[rootIndex1] == depth[rootIndex2]) && (rootIndex1 != rootIndex1) {
			depth[rootIndex1]++
		}
	}
}

func FindParentIndex(index int) int {
	if parentIndex[index] == index {
		return index
	} else {
		parentIndex[index] = FindParentIndex(parentIndex[index])
		return parentIndex[index]
	}
}

func SpanningTree() float64 {
	var pathLength float64
	pathVerticesCount := 0

	for i := 0; i < len(trails) && pathVerticesCount < verticesCount-1; i++ {
		trail := trails[i]
		if FindParentIndex(trail.FromIndex) != FindParentIndex(trail.ToIndex) {
			pathLength += math.Sqrt(float64(trail.Length))
			Union(trail.FromIndex, trail.ToIndex)
			pathVerticesCount++
		}
	}

	return pathLength
}

func main() {
	var x, y int
	fmt.Scanf("%d", &verticesCount)
	graph := make([]Vertex, verticesCount)
	depth = make([]int, verticesCount)
	parentIndex = make([]int, verticesCount)
	trails = make([]Trail, (verticesCount*(verticesCount+1))/2)

	for i := 0; i < verticesCount; i++ {
		fmt.Scanf("%d %d", &x, &y)
		graph[i] = Vertex{x, y}
		parentIndex[i] = i
	}

	counter := 0
	for i := 0; i < verticesCount; i++ {
		for j := i + 1; j < verticesCount; j++ {
			length := (graph[i].x-graph[j].x)*(graph[i].x-graph[j].x) + (graph[i].y-graph[j].y)*(graph[i].y-graph[j].y)
			trails[counter] = Trail{i, j, length}
			counter++
		}
	}

	sort.Slice(trails, func(i, j int) bool { return trails[i].Length < trails[j].Length })
	fmt.Printf("%.2f", SpanningTree())
}