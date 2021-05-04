// Based on this blog post:
// https://flaviocopes.com/golang-data-structure-graph/

package graph

type Vertex struct {
	Data interface{}
}

func NewVertex(data interface{}) Vertex {
	return Vertex{Data: data}
}

type Edge struct {
	Vertex Vertex
	Weight float64
}

type Graph struct {
	Vertices []Vertex
	Edges    map[Vertex][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make([]Vertex, 0),
		Edges:    make(map[Vertex][]Edge),
	}
}

func (g *Graph) AddVertex(v Vertex) {
	g.Vertices = append(g.Vertices, v)
}

func (g *Graph) AddEdge(v0, v1 Vertex, weight float64) {
	g.Edges[v0] = append(g.Edges[v0], Edge{v1, weight})
	g.Edges[v1] = append(g.Edges[v1], Edge{v0, weight})
}
