package main

import (
	"fmt"
)

const INFINITY = int(^uint(0) >> 1)

var (
	graph                     [][]int
	used                      []bool
	minEdgeLength             []int
	selectedEdge              []int
	verticesCount, roadsCount int
)

func PrimAlgorithm() int {
	result := 0
	for i := 0; i < verticesCount; i++ {
		newIndex := -1
		for j := 0; j < verticesCount; j++ {
			if !used[j] && (newIndex == -1 || minEdgeLength[j] < minEdgeLength[newIndex]) {
				newIndex = j
			}
		}
		used[newIndex] = true

		if selectedEdge[newIndex] != -1 {
			result += graph[newIndex][selectedEdge[newIndex]]
		}

		for toIndex := 0; toIndex < verticesCount; toIndex++ {
			if graph[newIndex][toIndex] < minEdgeLength[toIndex] {
				minEdgeLength[toIndex] = graph[newIndex][toIndex]
				selectedEdge[toIndex] = newIndex
			}
		}
	}
	return result
}

func main() {
	var roadV1, roadV2, roadLength int
	fmt.Scanf("%d\n%d", &verticesCount, &roadsCount)

	graph = make([][]int, verticesCount)
	used = make([]bool, verticesCount)
	minEdgeLength = make([]int, verticesCount)
	selectedEdge = make([]int, verticesCount)
	for i := range graph {
		minEdgeLength[i] = INFINITY
		selectedEdge[i] = -1
		graph[i] = make([]int, verticesCount)
		for j := range graph[i] {
			graph[i][j] = INFINITY
		}
	}

	for i := 0; i < roadsCount; i++ {
		fmt.Scanf("%d %d %d", &roadV1, &roadV2, &roadLength)
		graph[roadV1][roadV2] = roadLength
		graph[roadV2][roadV1] = roadLength
	}

	fmt.Println(PrimAlgorithm())
}
