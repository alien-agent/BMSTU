package main

import "fmt"

type Edge struct {
	firstIndex, secondIndex int
}

var (
	graph            [][]int
	visitedVertices  []bool
	currentComponent = *new([]int)
	maxComponent     = make(map[int]int)
	edges            []Edge
)

func componentEdgesCount(component []int) int {
	result := 0
	for _, vertexIndex := range component {
		result += len(graph[vertexIndex])
	}
	return result / 2
}

func componentBigger(comp1, comp2 []int) bool {
	if vCount1, vCount2 := len(comp1), len(comp2); vCount1 == vCount2 {
		if eCount1, eCount2 := componentEdgesCount(comp1), componentEdgesCount(comp2); eCount1 == eCount2 {
			// Т.к. ребра изначально отсортированы по возрастанию вершин и используется DFS,
			// вершины компонент также отсортированы по возрастанию. (минимальная - первая)
			return comp1[0] < comp2[0]
		} else {
			return eCount1 > eCount2
		}
	} else {
		return vCount1 > vCount2
	}
}

func depthFirstSearch(vertexIndex int) {
	visitedVertices[vertexIndex] = true
	currentComponent = append(currentComponent, vertexIndex)
	for edgeIndex := 0; edgeIndex < len(graph[vertexIndex]); edgeIndex++ {
		otherIndex := graph[vertexIndex][edgeIndex]
		if !visitedVertices[otherIndex] {
			depthFirstSearch(otherIndex)
		}
	}
}

func printGraph() {
	postfix := ""
	fmt.Println("graph {")
	// Print vertices
	for vertexIndex := range graph {
		if _, isRed := maxComponent[vertexIndex]; isRed {
			postfix = " [color = red]"
		} else {
			postfix = ""
		}
		fmt.Printf("\t%d%s\n", vertexIndex, postfix)
	}
	// Print edges
	for _, edge := range edges {
		if _, isRed := maxComponent[edge.firstIndex]; isRed {
			postfix = " [color = red]"
		} else {
			postfix = ""
		}
		fmt.Printf("\t%d -- %d%s\n", edge.firstIndex, edge.secondIndex, postfix)
	}
	fmt.Println("}")
}

func main() {
	var verticesCount, edgesCount int
	var v1, v2 int
	fmt.Scanf("%d\n%d", &verticesCount, &edgesCount)
	graph = make([][]int, verticesCount)
	edges = make([]Edge, edgesCount)
	visitedVertices = make([]bool, verticesCount)

	// Read graph
	for i := 0; i < verticesCount; i++ {
		graph[i] = *new([]int)
	}
	for i := 0; i < edgesCount; i++ {
		fmt.Scanf("%d %d", &v1, &v2)
		edges[i] = Edge{v1, v2}
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}

	// Find max component
	_maxComponent := *new([]int)
	for vertexIndex := range graph {
		if !visitedVertices[vertexIndex] {
			currentComponent = *new([]int)
			depthFirstSearch(vertexIndex)

			if len(_maxComponent) == 0 || componentBigger(currentComponent, _maxComponent) {
				_maxComponent = currentComponent
			}
		}
	}

	// Store max component as map[int]int to allow O(1) search
	for _, vertexIndex := range _maxComponent {
		maxComponent[vertexIndex] = vertexIndex
	}

	printGraph()
}
