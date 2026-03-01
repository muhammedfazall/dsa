package main

import "fmt"

//structure
type Graph struct{
	adjList map[int][]int //stores connection
	directed bool //same structure for both graph types
}

//constructor
func NewGraph(directed bool) *Graph {
	return &Graph{
		adjList: make(map[int][]int),
		directed: directed,
	}
}

//add edge
func (g *Graph) AddEdge(src, dest int)  {

	g.adjList[src] = append(g.adjList[src], dest)

	if !g.directed{
		g.adjList[dest] = append(g.adjList[dest], src)
	}
}

// to print
func (g *Graph) Print(){
	for vertex,neighbors := range g.adjList{
		fmt.Printf("%d -> %v\n",vertex,neighbors)
	}
}

//to create an isolated vertex
func (g *Graph) AddVertex(v int){
	if _,ok := g.adjList[v]; !ok{
		g.adjList[v] = []int{}
	}
}


func main()  {
	g:= NewGraph(false)

	g.AddEdge(0,1)
	g.AddEdge(0,2)
	g.AddEdge(1,2)
	g.AddEdge(2,3)
	g.adjList[5] = []int{}
	
	g.Print()
	
	fmt.Println("-----------------")
	dg:= NewGraph(true)
	
	dg.AddEdge(0,1)
	dg.AddEdge(0,2)
	dg.AddEdge(1,2)
	dg.AddEdge(2,3)
	dg.adjList[2] = append(dg.adjList[2], 5)

	dg.AddVertex(10)


	dg.Print()
}