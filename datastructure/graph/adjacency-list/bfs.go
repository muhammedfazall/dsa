package main

import "fmt"

func (g *Graph) BFS(start int) {

	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print(current, " ")
		
		for _,neighbor := range g.adjList[current]{
			if !visited[neighbor]{
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}