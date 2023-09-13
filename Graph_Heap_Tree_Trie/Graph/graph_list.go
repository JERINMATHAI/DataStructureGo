package main

import "fmt"

type graph struct {
	vertices []*vertex
}

type vertex struct {
	key      int
	adjacent []*vertex
}

func (g *graph) addVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Printf("Vertex %v already exist", k)
	} else {
		newNode := new(vertex)
		newNode.key = k
		g.vertices = append(g.vertices, newNode)
	}
}

func (g *graph) addEdges(from int, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("\nInvalid edge(%v-->%v)\n", from, to)
	} else if contains(fromVertex.adjacent, to) {
		fmt.Printf("\nExisting edge(%v-->%v)\n", from, to)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func contains(s []*vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *graph) getVertex(k int) *vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//print vertices and adjacent list
func (g *graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nvertices %v ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("\t%v", v.key)
		}
	}
}

func main() {
	fmt.Println("Graph - Adjacency list")
	myGraph := graph{}
	for i := 1; i < 5; i++ {
		myGraph.addVertex(i)
	}
	myGraph.addEdges(1, 2)
	myGraph.addEdges(1, 4)

	myGraph.addEdges(2, 1)

	myGraph.addEdges(3, 1)
	myGraph.addEdges(3, 4)
	myGraph.addEdges(3, 5)

	myGraph.addEdges(4, 2)
	myGraph.addEdges(4, 5)

	myGraph.Print()
}
