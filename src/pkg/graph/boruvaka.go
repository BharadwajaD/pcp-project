package graph

import (
	"context"
	"log"
	"sync"

	"github.com/theTechTrailBlazer/pcp-project/pkg/utils"
)

type boruvakaUtil struct {
	ufs       []utils.UnionFind
	nvertices int
}

func newboruvakaUtil(ntasks, nvertices int) boruvakaUtil {
	var ufs []utils.UnionFind

	for i := 0; i < ntasks; i++ {
		ufs = append(ufs, *utils.NewUnionFind(nvertices))
	}

	return boruvakaUtil{
		ufs:       ufs,
		nvertices: nvertices,
	}
}

func no_vetices(edgelist [][]int) int {
	vertices := make(map[int]bool)

	for _, edge := range edgelist {
		u, v := edge[0], edge[1]
		vertices[u] = true
		vertices[v] = true
	}

	return len(vertices)
}

func (b *boruvakaUtil) boruvakaLocal(edjlist [][]int, tid int) [][]int {
	uf := b.ufs[tid]
	defer uf.Reset()

	var mst [][]int
	var cheapest [][]int

	for i := 0; i < b.nvertices; i++ {
		cheapest = append(cheapest, []int{-1, -1, -1})
	}

	numComponents := no_vetices(edjlist)
    var prevComponents int
	mstWt := 0

	for numComponents > prevComponents {


		for _, edge := range edjlist {
			u := edge[0]
			v := edge[1]
			wt := edge[2]

			up := uf.Find(u)
			vp := uf.Find(v)

			if up != vp {
				if cheapest[up][2] == -1 || cheapest[up][2] > wt {
					cheapest[up] = []int{u, v, wt}
				}
				if cheapest[vp][2] == -1 || cheapest[vp][2] > wt {
					cheapest[vp] = []int{u, v, wt}
				}
			}
		}


		for i := 0; i < b.nvertices; i++ {
			if cheapest[i][2] != -1 {
				u := cheapest[i][0]
				v := cheapest[i][1]
				wt := cheapest[i][2]

				up := uf.Find(u)
				vp := uf.Find(v)

				if up != vp {
					mstWt += wt
					uf.Union(up, vp)
					mst = append(mst, []int{u, v, wt})
					numComponents--
				}
			}
		}

		//fmt.Printf("Thread %d: Finished component merging\n", tid)

		for i := 0; i < b.nvertices; i++ {
			cheapest[i][2] = -1
		}

        prevComponents = numComponents
	}

	return mst
}

func (g *GraphEdges) Boruvaka(ctx context.Context) [][]int {

	var ntasks int

	if tc, ok := ctx.Value("THREAD_COUNT").(int); ok {
		ntasks = tc
	} else {
		log.Fatalln("THREAD_COUNT is not int")
	}

	var wg sync.WaitGroup
	elist := g.EdgesList
	var elistMutex sync.Mutex
	var mstlist [][]int
	b := newboruvakaUtil(ntasks, g.NVertices)

	for ntasks > 0 {

		for tid := 0; tid < ntasks; tid++ {
			wg.Add(1)
			go func(tid int) {
				defer wg.Done()

				split := elist[tid*len(elist)/ntasks : (tid+1)*len(elist)/ntasks]
				localMST := b.boruvakaLocal(split, tid)

				// Use a mutex to protect concurrent append to the global elist
				elistMutex.Lock()
				mstlist = append(mstlist, localMST...)
				elistMutex.Unlock()
			}(tid)
		}

		wg.Wait()
		elist = mstlist
		mstlist = mstlist[:0]
		ntasks /= 2
	}

	return elist
}
