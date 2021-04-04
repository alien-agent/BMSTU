package main

import "fmt"

const INFINITY = int(^uint(0) >> 1)

var (
	graph                     [][]int
	visited                   []bool
	enterTime, fup            []int
	verticesCount, edgesCount int
	bridgesCount              = 0
	timer                     = 0
)

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func depthFirstSearch(vertexIndex, parentIndex int) {
	timer++
	visited[vertexIndex] = true
	enterTime[vertexIndex], fup[vertexIndex] = timer, timer

	for _, neighbourIndex := range graph[vertexIndex] {
		if neighbourIndex == parentIndex {
			continue
		}
		if visited[neighbourIndex] {
			fup[vertexIndex] = min(fup[vertexIndex], enterTime[neighbourIndex])
		} else {
			depthFirstSearch(neighbourIndex, vertexIndex)
			fup[vertexIndex] = min(fup[vertexIndex], fup[neighbourIndex])
			if fup[neighbourIndex] > enterTime[vertexIndex] {
				bridgesCount++
			}
		}
	}
}

func main() {
	var v1, v2 int
	fmt.Scanf("%d\n%d", &verticesCount, &edgesCount)
	// Init arrays
	graph = make([][]int, verticesCount)
	visited = make([]bool, verticesCount)
	enterTime = make([]int, verticesCount)
	fup = make([]int, verticesCount)
	for i := 0; i < verticesCount; i++ {
		graph[i] = *new([]int)
		enterTime[i], fup[i] = INFINITY, INFINITY
	}
	// Fill graph
	for i := 0; i < edgesCount; i++ {
		fmt.Scanf("%d %d", &v1, &v2)
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}
	// DFS to find bridges
	for i := 0; i < verticesCount; i++ {
		if !visited[i] {
			depthFirstSearch(i, -1)
		}
	}

	fmt.Println(bridgesCount)
}
