package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/theTechTrailBlazer/pcp-project/pkg/graph"
)

func get_input(input_file string) string {
	var input string
	if input_file == "" {
		input, err := os.Getwd()
		if err != nil {
            log.Fatal("get_input:1:", err)
		}

		input = filepath.Join(input, "datasets", "input.txt")
		log.Println(input)
	} else {
		input = input_file
	}

	data, err := os.ReadFile(input)
	if err != nil {
        log.Fatal("get_input:2:", err)
	}

	input = string(data)

	return input
}

func write_output(output_file string, content *string) {
}

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    thread_count := os.Getenv("THREAD_COUNT")

	input_file := flag.String("input", "", "input dataset file")
	//ouput_file := flag.String("output", "", "output file")
	flag.Parse()

    ctx := context.Background()
    ctx = context.WithValue(ctx, "THREAD_COUNT", thread_count)


	input := get_input(*input_file)
	graph_input := strings.Split(input, "\n")
	_ = graph.NewGraphEdges(&graph_input)

	gadj := graph.NewGraphAdj(&graph_input)
	gadj.Boruvaka(ctx)

}
