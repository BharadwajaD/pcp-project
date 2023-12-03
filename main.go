package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	//ouput_file := flag.String("output", "", "output file")
	flag.Parse()

    ctx := context.Background()
    ctx = context.WithValue(ctx, "THREAD_COUNT", thread_count)


	input := get_input(*input_file)
	graph_input := strings.Split(input, "\n")

    gedj := graph.NewGraphEdgesString(&graph_input)
    gedj.Boruvaka(ctx)


}
