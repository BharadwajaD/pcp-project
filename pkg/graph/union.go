package graph

// UnionFind represents the Union-Find (Disjoint Set) data structure.
type UnionFind struct {
	root []int
	rank []int
}

// NewUnionFind creates a new Union-Find data structure with the given number of elements.
func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		root: make([]int, size),
		rank: make([]int, size),
	}

	for i := 0; i < size; i++ {
		uf.root[i] = i // Initialize each element as its own parent
		uf.rank[i] = 0 // Initialize the rank of each element as 0
	}

	return uf
}

// Find returns the representative (root) of the set that contains the given element.
func (uf *UnionFind) Find(element int) int {
	if uf.root[element] != element {
		uf.root[element] = uf.Find(uf.root[element]) // Path Compression
	}
	return uf.root[element]
}

// Union combines the sets containing the given two elements.
func (uf *UnionFind) Union(element1, element2 int) {
	root1 := uf.Find(element1)
	root2 := uf.Find(element2)

	if root1 != root2 {
		// Union by rank: Attach the smaller tree to the root of the larger tree.
		if uf.rank[root1] < uf.rank[root2] {
			uf.root[root1] = root2
		} else if uf.rank[root1] > uf.rank[root2] {
			uf.root[root2] = root1
		} else {
			uf.root[root2] = root1
			uf.rank[root1]++
		}
	}
}

// Connected checks if the given two elements are in the same set.
func (uf *UnionFind) Connected(element1, element2 int) bool {
	return uf.Find(element1) == uf.Find(element2)
}
