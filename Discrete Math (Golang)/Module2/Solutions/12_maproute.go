package main

import "fmt"

const INFINITY = int(^uint(0) >> 1)

var (
	graph = make(map[int][]int)
	weights = make(map[int]int)
	distances = make(map[int]int)
	visited = make(map[int]bool)
)

func DijkstraComputeDistances() {
	for range graph {
		closestVertexIndex := -1
		for vertexIndex := range graph {
			if !visited[vertexIndex] && (closestVertexIndex == -1 || distances[vertexIndex] < distances[closestVertexIndex]) {
				closestVertexIndex = vertexIndex
			}
		}
		visited[closestVertexIndex] = true

		for _, neighbourIndex := range graph[closestVertexIndex]{
			newDistance := distances[closestVertexIndex] + weights[neighbourIndex]
			if newDistance < distances[neighbourIndex] {
				distances[neighbourIndex] = newDistance
			}
		}
	}
}

func main() {
	var mapSize, temp int
	fmt.Scanf("%d", &mapSize)

	getPossibleTrails := func(index int) []int {
		result := make([]int, 0)
		row, col := index/mapSize, index%mapSize
		if col+1 < mapSize {
			result = append(result, row*mapSize+(col+1))
		}
		if row+1 < mapSize {
			result = append(result, (row+1)*mapSize+col)
		}
		return result
	}

	for i := 0; i < mapSize*mapSize; i++ {
		fmt.Scanf("%d", &temp)
		graph[i] = getPossibleTrails(i)
		weights[i] = temp
		distances[i] = INFINITY
		visited[i] = false
	}

	distances[0] = weights[0]
	DijkstraComputeDistances()
	fmt.Println(distances[len(graph)-1])
}
