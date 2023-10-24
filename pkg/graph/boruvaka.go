package graph

import (
	"context"
	"log"
)

//need a way to pass context info like: no of threads, log file
func (g *GraphAdj)Boruvaka(ctx context.Context){
    log.Println("Boruvaka algo started")
    log.Println("No of threads: ", ctx.Value("THREAD_COUNT"))
    log.Println("Boruvaka algo ended")
}
