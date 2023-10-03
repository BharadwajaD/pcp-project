package graph

import (
	"strconv"
	"strings"
)

const (
    EDGES_VALUES = 3
)

//This is used for kruskal's and other algos
type GraphAdj struct {
    nodes_count int
    adj_list [][][EDGES_VALUES]int
}

func NewGraphAdj(input *[]string) GraphAdj {
	return GraphAdj{}
}

//This is used for prim's algo
type GraphEdges struct {
    nodes_count int
    edges_list [][EDGES_VALUES]int
}

func NewGraphEdges(input *[]string) GraphEdges{
    var edges_list [][EDGES_VALUES]int
    input_ := *input
    for _, e := range input_[:len(*input)-1]{
        val := strings.Split(e, " ")
        u, _ := strconv.Atoi(val[0]);
        v, _ := strconv.Atoi(val[1]);
        wt, _ := strconv.Atoi(val[2]);

        edges_list = append(edges_list, [3]int{u, v, wt})
    }

    return GraphEdges{
        nodes_count: len(*input),
        edges_list: edges_list,
    }
}
