package main

import "fmt"

// reachable only
func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	g.dfsHelper(start, visited)

}

func (g *Graph) dfsHelper(node int, visited map[int]bool) {
	visited[node] = true
	fmt.Print(node, " ")

	for _,neighbor := range g.adjList[node]{
		if !visited[neighbor]{
			g.dfsHelper(neighbor,visited)
		}
	}
}

//Start DFS from every unvisited node (entire graph)
func (g *Graph) DFSALL(){
	visited := make(map[int]bool)

	for vertex := range g.adjList{
		if !visited[vertex]{
			g.dfsHelper(vertex,visited)
		}
	}
}