package graph

import (
	"context"
	"log"
	"sync"
)


type primsUtil struct{
    nvertices int
    ntasks int
    key []int
    visited []int
}

func newprimsUtil(nvertices, ntasks int) primsUtil {
    return primsUtil{
        nvertices: nvertices,
        ntasks: ntasks,
        key:  make([]int, nvertices),
        visited: make([]int, nvertices),
    }
}

func (p *primsUtil) minKey(wg *sync.WaitGroup, mu *sync.Mutex) int {
	defer wg.Done()

	min := int(^uint(0) >> 1)
	var index, i int

	for i = 0; i < p.nvertices; i++ {
		if p.visited[i] == 0 && p.key[i] < min {
			min = p.key[i]
			index = i
		}
	}

	mu.Lock()
	if min < p.key[p.ntasks] {
		p.ntasks = index
	}
	mu.Unlock()

	return index
}

func (g *GraphAdjMatrix) Prims(ctx context.Context) [][]int {

	var ntasks int

	if tc, ok := ctx.Value("THREAD_COUNT").(int); ok {
		ntasks = tc
	} else {
		log.Fatalln("THREAD_COUNT is not int")
	}

	graph := g.AdjMatrix
    p := newprimsUtil(g.Nvertices, ntasks)

	from := make([]int, p.nvertices)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < p.nvertices; i++ {
		p.key[i] = int(^uint(0) >> 1)
		p.visited[i] = 0
	}

	p.key[0] = 0
	from[0] = -1

	for count := 0; count < p.nvertices-1; count++ {
		u := p.minKey(&wg, &mu)
		p.visited[u] = 1

		var v int
		for v = 0; v < p.nvertices; v++ {
			if graph[u][v] != 0 && p.visited[v] == 0 && graph[u][v] < p.key[v] {
				from[v] = u
				p.key[v] = graph[u][v]
			}
		}
	}

    return make([][]int, 0)
}
