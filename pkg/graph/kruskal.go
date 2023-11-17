package graph

import (
	"context"
	"log"
)

//need a way to pass context info like: no of threads, log file
func (g *GraphEdges)Kruskal(ctx context.Context){
    log.Println("Kruskal: No of threads: ", ctx.Value("THREAD_COUNT"))
}
