package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/theTechTrailBlazer/pcp-project/pkg/graph"
)

func get_input(input_file string) string {
	var input string
	if input_file == "" {

		input = filepath.Join("datasets", "input.txt")
	} else {
		input = input_file
	}

	data, err := os.ReadFile(input)
	if err != nil {
        log.Fatal(err)
	}

	input = string(data)

	return input
}

// PrintMST prints the Minimum Spanning Tree (MST).
func PrintMST(mstEdges [][]int) {
	fmt.Println("Minimum Spanning Tree (MST):")
	for _, edge := range mstEdges {
		u, v, wt := edge[0], edge[1], edge[2]
		fmt.Printf("%d - %d : %d\n", u, v, wt)
	}
}

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    thread_count, err := strconv.Atoi(os.Getenv("THREAD_COUNT"))
    if err != nil {
		log.Fatalln("THREAD_COUNT is not int")
    }

	input_file := flag.String("input", "", "input dataset file")
	algo := flag.String("algo", "kruskal", "alogrithm to construct mst")
	flag.Parse()

    ctx := context.Background()
    ctx = context.WithValue(ctx, "THREAD_COUNT", thread_count)


	input := get_input(*input_file)
	graph_input := strings.Split(input, "\n")
    nvertices,_ := strconv.Atoi(graph_input[0])
    graph_input = graph_input[1:]

    graphEdj := graph.NewGraphEdges(nvertices, &graph_input)
    graphAdjMatrix := graph.NewGraphAdjMatrix(nvertices, &graph_input)

    var mst [][]int

    start_time := time.Now()

    if *algo == "prims" {
        mst = graphAdjMatrix.Prims(ctx)
    }else if *algo == "boruvaka"{
        mst = graphEdj.Boruvaka(ctx)
    }else{
        mst = graphEdj.Kruskal(ctx)
    }

    end_time := time.Now()

    _ = mst
    //PrintMST(mst)
    log.Printf("Time taken by %s algo is %v\n", *algo, end_time.Sub(start_time))
}
