package graph

import (
	"context"
	"log"
)

//need a way to pass context info like: no of threads, log file
func (g *GraphAdj)Prims(ctx context.Context){
    log.Println("Prims: No of threads: ", ctx.Value("THREAD_COUNT"))
}
