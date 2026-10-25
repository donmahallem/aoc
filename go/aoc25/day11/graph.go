package day11

import (
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

// graph stores nodes as dense ids with their successors in compressed sparse row form.
type graph struct {
	ids     *nodeIndex
	offsets []int32
	edges   []int32
}

// hashMul is the Fibonacci hashing multiplier used to spread packed keys.
const hashMul = 0x9E3779B97F4A7C15

// nodeIndex maps packed node keys to dense ids using open addressing with linear probing.
type nodeIndex struct {
	slots []nodeSlot
	count int32
}

// nodeSlot pairs a key with its id; key 0 marks an empty slot since encodeKey never returns 0.
type nodeSlot struct {
	key uint32
	id  int32
}

func newNodeIndex(expected int) *nodeIndex {
	size := 16
	for size < expected*2 {
		size <<= 1
	}
	return &nodeIndex{slots: make([]nodeSlot, size)}
}

func (t *nodeIndex) size() int {
	return int(t.count)
}

func (t *nodeIndex) find(key uint32) (int32, bool) {
	mask := uint64(len(t.slots) - 1)
	for i := ((uint64(key) * hashMul) >> 32) & mask; ; i = (i + 1) & mask {
		s := &t.slots[i]
		if s.key == key {
			return s.id, true
		}
		if s.key == 0 {
			return 0, false
		}
	}
}

// intern returns the id of key, assigning the next id if key is new.
func (t *nodeIndex) intern(key uint32) int32 {
	if int(t.count)*2 >= len(t.slots) {
		t.grow()
	}
	mask := uint64(len(t.slots) - 1)
	for i := ((uint64(key) * hashMul) >> 32) & mask; ; i = (i + 1) & mask {
		s := &t.slots[i]
		if s.key == key {
			return s.id
		}
		if s.key == 0 {
			s.key = key
			s.id = t.count
			t.count++
			return s.id
		}
	}
}

func (t *nodeIndex) grow() {
	old := t.slots
	t.slots = make([]nodeSlot, 2*len(old))
	mask := uint64(len(t.slots) - 1)
	for _, s := range old {
		if s.key == 0 {
			continue
		}
		j := ((uint64(s.key) * hashMul) >> 32) & mask
		for t.slots[j].key != 0 {
			j = (j + 1) & mask
		}
		t.slots[j] = s
	}
}

// buildGraph converts the flat (from, to) id pairs into a graph.
func buildGraph(ids *nodeIndex, pairs []int32) *graph {
	n := ids.size()
	g := &graph{
		ids:     ids,
		offsets: make([]int32, n+1),
		edges:   make([]int32, len(pairs)/2),
	}
	for i := 0; i < len(pairs); i += 2 {
		g.offsets[pairs[i]+1]++
	}
	for i := 1; i <= n; i++ {
		g.offsets[i] += g.offsets[i-1]
	}
	// Advancing offsets[from] per edge leaves offsets[i] at the start of i+1, so shift right afterwards.
	for i := 0; i < len(pairs); i += 2 {
		from := pairs[i]
		g.edges[g.offsets[from]] = pairs[i+1]
		g.offsets[from]++
	}
	copy(g.offsets[1:], g.offsets[:n])
	g.offsets[0] = 0
	return g
}

func (g *graph) successors(id int32) []int32 {
	return g.edges[g.offsets[id]:g.offsets[id+1]]
}

// pathCounter counts paths in a graph, reusing its buffers across calls.
type pathCounter struct {
	graph *graph
	order []int32
	pos   []int32
	ways  []uint64
}

// newPathCounter returns a SolverError if the graph contains a cycle.
func newPathCounter(g *graph) (*pathCounter, error) {
	n := g.ids.size()
	indegree := make([]int32, n)
	for _, next := range g.edges {
		indegree[next]++
	}

	order := make([]int32, 0, n)
	for id := range n {
		if indegree[id] == 0 {
			order = append(order, int32(id))
		}
	}
	for head := 0; head < len(order); head++ {
		for _, next := range g.successors(order[head]) {
			indegree[next]--
			if indegree[next] == 0 {
				order = append(order, next)
			}
		}
	}
	if len(order) != n {
		return nil, aoc_utils.NewSolverErrorf("graph contains a cycle")
	}
	// indegree is zero after a complete sort, so its storage is reused for pos.
	pos := indegree
	for i, id := range order {
		pos[id] = int32(i)
	}
	return &pathCounter{graph: g, order: order, pos: pos, ways: make([]uint64, n)}, nil
}

// count returns the number of paths from start to target that never enter a skipped key.
func (p *pathCounter) count(start, target uint32, skip ...uint32) uint64 {
	if start == target {
		return 1
	}
	startID, okStart := p.graph.ids.find(start)
	targetID, okTarget := p.graph.ids.find(target)
	if !okStart || !okTarget {
		return 0
	}

	var blockedBuf [4]int32
	blocked := blockedBuf[:0]
	for _, key := range skip {
		if id, ok := p.graph.ids.find(key); ok {
			blocked = append(blocked, id)
		}
	}
	if slices.Contains(blocked, startID) {
		return 0
	}

	startPos, targetPos := p.pos[startID], p.pos[targetID]
	if startPos > targetPos {
		return 0
	}

	clear(p.ways)
	p.ways[startID] = 1
	for _, id := range p.order[startPos:targetPos] {
		paths := p.ways[id]
		if paths == 0 {
			continue
		}
		for _, next := range p.graph.successors(id) {
			if !slices.Contains(blocked, next) {
				p.ways[next] += paths
			}
		}
	}
	return p.ways[targetID]
}
