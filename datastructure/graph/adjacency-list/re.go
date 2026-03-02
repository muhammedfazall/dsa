package main  //revisions

// import "fmt"

// type Graph struct {
// 	adjL     map[int][]int
// 	directed bool
// }

// func NewGraph(directed bool) *Graph {
// 	return &Graph{
// 		adjL:     make(map[int][]int),
// 		directed: directed,
// 	}
// }

// func (g *Graph) AddEdge(src, dest int) {
// 	g.adjL[src] = append(g.adjL[src], dest)
// 	if !g.directed {
// 		g.adjL[dest] = append(g.adjL[dest], src)
// 	}
// }

// func (g *Graph) Print() {
// 	for vertex, neighbor := range g.adjL {
// 		fmt.Printf("%d -> %v",vertex,neighbor)
// 	}
// }

// func (g *Graph) AddVertex(v int)  {
// 	if _,ok := g.adjL[v] ; !ok{
// 		g.adjL[v] = []int{}
// 	}
// }

// func (g *Graph) BFS(start int){
// 	visited := make(map[int]bool)
// 	queue := []int{start}

// 	visited[start] = true

// 	for len(queue) > 0 {
// 		cur := queue[0]
// 		queue = queue[1:]

// 		fmt.Print(cur, " ")

// 		for _,nieghbor := range g.adjL[cur]{
// 			if !visited[nieghbor]{
// 				visited[nieghbor] = true
// 				queue = append(queue, nieghbor)
// 			}
// 		}
// 	}
// }

// func (g *Graph) DFS(start int){
// 	visited := make(map[int]bool)
// 		stack := []int{start}

// 	visited[start] = true

// 	for len(stack) > 0 {
// 		cur := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		fmt.Print(cur, " ")

// 		for _,nieghbor := range g.adjL[cur]{
// 			if !visited[nieghbor]{
// 				visited[nieghbor] = true
// 				stack = append(stack, nieghbor)
// 			}
// 		}
// 	}
// }

// func (g *Graph) DFS(start int)  {
// 	visited := make(map[int]bool)
// 	g.dfsHelper(start,visited)
// }

// func (g *Graph) dfsHelper(node int, visited map[int]bool)  {
// 	visited[node] = true
// 	fmt.Print(node, " ")

// 	for _,neighbor := range g.adjL[node]{
// 		if !visited[neighbor]{
// 			g.dfsHelper(neighbor,visited)
// 		}
// 	}
// }

// func (g *Graph) DFSALL()  {
// 	visited := make(map[int]bool)

// 	for vertex := range g.adjL{
// 		if !visited[vertex]{
// 			g.dfsHelper(vertex,visited)
// 		}
// 	}
// }