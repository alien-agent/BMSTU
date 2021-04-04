package main

import (
	"container/list"
	"fmt"
)

var (
	graph              [][]int
	verticesCount, edgesCount int
)

var (
	pillows            []int
	distances          [][]int // distances[i][j] - расстояние от i-той вершины до j-той опоры
	visited            map[int]bool
)

type Node struct {
	vertexIndex, depth int
}

func (n Node) AsTuple() (int, int){
	return n.vertexIndex, n.depth
}

func readGraph(){
	var v1, v2 int
	fmt.Scanf("%d\n%d", &verticesCount, &edgesCount)
	graph = make([][]int, verticesCount)

	for i := 0; i < verticesCount; i++ {
		graph[i] = *new([]int)
	}
	for i := 0; i < edgesCount; i++ {
		fmt.Scanf("%d %d", &v1, &v2)
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}
}

func readPillows(){
	var pillowsCount int
	fmt.Scanf("%d", &pillowsCount)
	if pillowsCount == 0{
		fmt.Scanf("%d", &pillowsCount)
	}
	pillows = make([]int, pillowsCount)

	for i := 0; i < pillowsCount; i++ {
		fmt.Scanf("%d", &pillows[i])
	}

	for i := 0; i < verticesCount; i++ {
		distances[i] = make([]int, pillowsCount)
	}
}

func breadthFirstSearch(startIndex, currentPillowIndex int) {
	queue := list.New()
	queue.PushBack(Node{startIndex, 0})
	visited[startIndex] = true

	for queue.Len() > 0 {
		qnode := queue.Front()
		vertexIndex, currentDepth := qnode.Value.(Node).AsTuple()
		for _, neighbourIndex := range graph[vertexIndex] {
			if !visited[neighbourIndex]{
				visited[neighbourIndex] = true
				distances[neighbourIndex][currentPillowIndex] = currentDepth + 1
				queue.PushBack(Node{neighbourIndex, currentDepth + 1})
			}
		}
		queue.Remove(qnode)
	}
}

func sliceIsDuplicatesOnly(slice []int) bool {
	if len(slice) < 2 {
		return true
	}

	initial := slice[0]
	for _, val := range slice {
		if val != initial {
			return false
		}
	}
	return true
}

func computeDistances() {
	for pillowIndex, pillowIndexInGraph := range pillows {
		visited = make(map[int]bool)
		breadthFirstSearch(pillowIndexInGraph, pillowIndex)
	}
}

func computeResult() []int {
	result := *new([]int)
	for vertexIndex, vertexDistances := range distances {
		// Если существует только 1 опора, то ее массив растояний будет
		// равен [0] (расстояние до себя), и она ошибочно попадет в ответ.
		if sliceIsDuplicatesOnly(vertexDistances) && vertexDistances[0] != 0 {
			result = append(result, vertexIndex)
		}
	}
	return result
}

func main() {
	readGraph()
	distances = make([][]int, verticesCount)
	readPillows()

	computeDistances()
	result := computeResult()

	if len(result) == 0 {
		fmt.Println("-")
	} else {
		for _, vertexIndex := range result {
			fmt.Printf("%d ", vertexIndex)
		}
	}
}
