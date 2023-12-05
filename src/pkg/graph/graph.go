package graph

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	EDGES_VALUES = 3
	PAIR_VALUES = 2
)

// This is used for kruskal's and other algos
type GraphAdj struct {
	Nvertices int
	AdjList    [][][]int
}

func NewGraphAdj(nvertices int, input *[]string) GraphAdj {

    adj_list := make([][][]int, nvertices)

	input_ := *input
	for _, e := range input_[:len(*input)-1] {
		val := strings.Split(e, " ")
		u, _ := strconv.Atoi(val[0])
		v, _ := strconv.Atoi(val[1])
		wt, _ := strconv.Atoi(val[2])

		adj_list[u-1] = append(adj_list[u-1], []int{v-1, wt})
		adj_list[v-1] = append(adj_list[v-1], []int{u-1, wt})
	}

	return GraphAdj{
        Nvertices: nvertices,
        AdjList: adj_list,
	}

}

// PrintGraph prints the graph in the format "u : v, wt"
func (g *GraphAdj) PrintGraph() {
	for u, edges := range g.AdjList {
        fmt.Printf("%d\n", u)
		for _, edge := range edges {
			v, wt := edge[0], edge[1]
            fmt.Printf("[%d, %d]",  v, wt)
		}
        fmt.Printf("\n")
	}
}

// This is used for prim's algo
type GraphEdges struct {
	NVertices int
	EdgesList [][]int
}

func NewGraphEdges(nvertices int, input *[]string) GraphEdges {
	var edges_list [][]int
	input_ := *input
	for _, e := range input_[:len(*input)-1] {
		val := strings.Split(e, " ")
		u, _ := strconv.Atoi(val[0])
		v, _ := strconv.Atoi(val[1])
		wt, _ := strconv.Atoi(val[2])

		edges_list = append(edges_list, []int{u-1, v-1, wt})
	}

	return GraphEdges{
		NVertices: nvertices,
		EdgesList: edges_list,
	}
}

// PrintGraph prints the graph in the format "u : v, wt"
func (g *GraphEdges) PrintGraph() {
	for _, edge := range g.EdgesList {
        fmt.Println(edge)
	}
}

type GraphAdjMatrix struct{
    Nvertices int
    AdjMatrix [][]int
}

func NewGraphAdjMatrix(nvertices int, input *[]string) GraphAdjMatrix {
    var adjmatrix [][]int

    for i := 0; i < nvertices; i++ {
        adjmatrix = append(adjmatrix, make([]int, nvertices))
    }

	input_ := *input
	for _, e := range input_[:len(*input)-1] {
		val := strings.Split(e, " ")
		u, _ := strconv.Atoi(val[0])
		v, _ := strconv.Atoi(val[1])
		wt, _ := strconv.Atoi(val[2])

        adjmatrix[u-1][v-1] = wt
	}

    return GraphAdjMatrix{
        Nvertices: nvertices,
        AdjMatrix: adjmatrix,
    }
}
