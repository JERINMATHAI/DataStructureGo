package main

import "fmt"

const vertices = 5

type graph struct {
	adjMatrix [5][5]int
}

// Initialize adjacency matrix
func (g *graph) init() [5][5]int {
	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			g.adjMatrix[i][j] = 0
		}
	}
	return g.adjMatrix
}

//add edges
func (g *graph) addEdge(i int, j int) {
	g.adjMatrix[i][j] = 1
}

//print graph
func (g *graph) printGraph() {
	for i := 0; i < vertices; i++ {
		if i != 0 {
			fmt.Printf("%d", i)
		}
		for j := 0; j < vertices; j++ {
			if i == 0 && j == 0 {
				fmt.Println(" ")
				for k := 0; k < vertices; k++ {
					fmt.Printf("%d", k)
				}
				fmt.Println(" ")
				fmt.Printf("%d", i)
			}
			fmt.Printf("%d", g.adjMatrix[i][j])
		}
		fmt.Println(" ")
	}
}

func main() {
	fmt.Println("Graph using matrix")

	mygraph := graph{}
	mygraph.init()
	mygraph.addEdge(4, 1)
	mygraph.addEdge(1, 2)
	mygraph.addEdge(4, 3)
	mygraph.addEdge(4, 2)
	mygraph.addEdge(1, 4)
	mygraph.addEdge(2, 4)
	mygraph.addEdge(2, 1)
	mygraph.printGraph()

}
