package main

import "fmt"

type Graph struct{
	al map[int][]int
	directed bool
}

func MakeGraph(directed bool) *Graph {
	return &Graph{
		al: make(map[int][]int),
		directed: directed,
	}
}

func (g *Graph) AddEdge(src,dest int){
	g.al[src] = append(g.al[src],dest)
}

func (g *Graph) BSF(start int){
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		fmt.Print(cur, " ")

		for _,n := range g.al[cur] {
			if !visited[n]{
				queue = append(queue, n)
			} 
		}
	}
}

// func main()  {
// 	g := MakeGraph(false)

// 	g.AddEdge(0,10)
// 	g.AddEdge(0,1)

// 	g.BSF(0)
// }