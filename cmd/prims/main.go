package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/theTechTrailBlazer/pcp-project/pkg/graph"
)

func get_input(input_file string) string {
	var input string
	if input_file == "" {
		input, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		input = filepath.Join(input, "datasets", "input.txt")
	} else {
		input = input_file
	}

	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	input = string(data)

	return input
}

func write_output(output_file string, content *string){
}

func main() {

	input_file := flag.String("input", "", "input dataset file")
	//ouput_file := flag.String("output", "", "output file")

    flag.Parse()

    input := get_input(*input_file)
    graph_input := strings.Split(input, "\n");
    _ = graph.NewGraphEdges(&graph_input)

}
