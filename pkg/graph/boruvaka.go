package graph

import (
	"context"
	"log"
	"sort"
	"sync"
)

type boruvaka struct {
	graph *GraphEdges
	mst   GraphEdges
	mutex sync.Mutex
	uf    *UnionFind
}

func newBoruvaka(g *GraphEdges) boruvaka {
	uf := NewUnionFind(g.NVertices)

	return boruvaka{
		uf:    uf,
		graph: g,
		mst:   NewGraphEdges(),
	}
}

func (b *boruvaka) runParallel(tid int, edges [][EDGES_VALUES]int) {
	//1. Sort the edges based on the weight
	//2. Merge the components and update MST

	log.Printf("runParallel %d", tid)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		b.mutex.Lock()

		c1 := b.uf.Find(u)
		c2 := b.uf.Find(v)

		if c1 != c2 {
			b.uf.Union(c1, c2)
			b.mst.EdgesList = append(b.mst.EdgesList, edge)
		}

		b.mutex.Unlock()
	}
}

// return a new graph which is a MST of the given graph
func (g *GraphEdges) Boruvaka(ctx context.Context) *GraphEdges {

	var ntasks int

	if tc, ok := ctx.Value("THREAD_COUNT").(int); ok {
		ntasks = tc
	} else {
		log.Fatalln("THREAD_COUNT is not int")
	}

	b := newBoruvaka(g)

	//here we are creating n tasks...
	//Note: No of tasks can be greater than number of threads
	//Here you are putting constraint on number of tasks and not on number of threads

	edges := g.EdgesList
    nedges := len(edges)

	var wg sync.WaitGroup
	for tid := 0; tid < ntasks; tid++ {
		wg.Add(1)
		go func(tid int) {
			defer wg.Done()
			b.runParallel(tid, edges[tid*nedges/ntasks:(tid+1)*nedges/ntasks])

		}(tid)
	}

	wg.Wait()

    log.Printf("Boruvaka mst len: %d and graph len %d", len(b.mst.EdgesList), len(g.EdgesList))

	return nil
}
