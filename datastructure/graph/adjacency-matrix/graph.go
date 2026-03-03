package main

import "fmt"

type GraphMatrix struct {
	matrix   [][]int
	vertices int
	directed bool
}

func NewGraphMatrix(vertices int, directed bool) *GraphMatrix {
	matrix := make([][]int, vertices)

	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}

	return &GraphMatrix{
		matrix:   matrix,
		vertices: vertices,
		directed: directed,
	}
}

func (g *GraphMatrix) AddEdge(src, dest int) {
	g.matrix[src][dest] = 1

	if !g.directed {
		g.matrix[dest][src] = 1
	}
}

func (g *GraphMatrix) Print() {
	for _, row := range g.matrix {
		fmt.Println(row)
	}
}
