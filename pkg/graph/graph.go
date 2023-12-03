package graph

import (
	"strconv"
	"strings"
)

const (
	EDGES_VALUES = 3
)

// This is used for kruskal's and other algos
type GraphAdj struct {
	nodes_count int
	adj_list    [][][EDGES_VALUES]int
}

func NewGraphAdj(input *[]string) GraphAdj {
	return GraphAdj{}
}

// This is used for prim's algo
type GraphEdges struct {
	NVertices int
	EdgesList [][EDGES_VALUES]int
}

func NewGraphEdges() GraphEdges {
	return GraphEdges{NVertices: 0, EdgesList: make([][3]int, 0)}
}

func NewGraphEdgesString(input *[]string) GraphEdges {
	var edges_list [][EDGES_VALUES]int
	input_ := *input
	for _, e := range input_[:len(*input)-1] {
		val := strings.Split(e, " ")
		u, _ := strconv.Atoi(val[0])
		v, _ := strconv.Atoi(val[1])
		wt, _ := strconv.Atoi(val[2])

		edges_list = append(edges_list, [3]int{u, v, wt})
	}

	return GraphEdges{
		NVertices: len(*input),
		EdgesList: edges_list,
	}
}
