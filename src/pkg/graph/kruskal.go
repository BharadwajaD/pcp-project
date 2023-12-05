package graph

import (
	"context"
	"log"
	"sort"
	"sync"

	"github.com/theTechTrailBlazer/pcp-project/pkg/utils"
)



type kruskalUtil struct{
    ufs []utils.UnionFind;
}

func newKruskalUtil(ntasks, nvertices int) kruskalUtil {
    var ufs []utils.UnionFind

    for i := 0; i < ntasks; i++ {
        ufs = append(ufs, *utils.NewUnionFind(nvertices))
    }

    return kruskalUtil{
        ufs: ufs,
    }
}

func split_edge_list(edjlist [][]int, nsplits int) [][][]int {
    var splits [][][]int
    nedges := len(edjlist)

    for i:= 0; i < nsplits; i++ {
        splits = append(splits, edjlist[i*nedges/nsplits: (i+1)*nedges/nsplits])
    }

    return splits
}

func (k *kruskalUtil)mstLocal(edjlist [][]int, nvertices int, tid int) [][]int {

	sort.Slice(edjlist, func(i, j int) bool {
		return edjlist[i][2] < edjlist[j][2]
	})

    uf := k.ufs[tid]
    defer uf.Reset()

	var mstEdges [][]int

	for _, edge := range edjlist {
		u := edge[0]
		v := edge[1]

		if uf.Find(u) != uf.Find(v) {
			uf.Union(u, v)
			mstEdges = append(mstEdges, edge)
		}
	}

	return mstEdges
}

func (g *GraphEdges) Kruskal(ctx context.Context) [][]int {
	var ntasks int

	if tc, ok := ctx.Value("THREAD_COUNT").(int); ok {
		ntasks = tc
	} else {
		log.Fatalln("THREAD_COUNT is not int")
	}

	var wg sync.WaitGroup
	elist := g.EdgesList
	var elistMutex sync.Mutex
    k := newKruskalUtil(ntasks, g.NVertices)

	for ntasks > 0 {
		splits := split_edge_list(elist, ntasks)
		elist = elist[:0]

		for tid := 0; tid < ntasks; tid++ {
			wg.Add(1)
			go func(tid int) {
				defer wg.Done()

				localElist := k.mstLocal(splits[tid], g.NVertices, tid)

				// Use a mutex to protect concurrent append to the global elist
				elistMutex.Lock()
				elist = append(elist, localElist...)
				elistMutex.Unlock()
			}(tid)
		}

		wg.Wait()
		ntasks /= 2
	}

	return elist
}

