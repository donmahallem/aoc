package day23

import (
	"fmt"
	"io"
)

type NodeHashList = []NodeHash

func StringifySequence(a []NodeHash) string {
	var out string = ""
	for _, item := range a {
		if len(out) > 0 {
			out += ","
		}
		unhashed := UnhashId(item)
		out += fmt.Sprintf("%c%c", unhashed[0], unhashed[1])
	}
	return out
}

func findLongest(data CombinationMap) []NodeHash {
	var longest []NodeHash
	n := len(data)
	path := make([]NodeHash, n)
	used := make(map[NodeHash]struct{}, n)

	var dfs func(depth int)
	dfs = func(depth int) {
		last := path[depth-1]
		extended := false
		for next := range data[last] {
			if _, ok := used[next]; ok {
				continue
			}
			// Check clique property
			valid := true
			for i := range depth {
				if _, ok := data[path[i]][next]; !ok {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}
			used[next] = struct{}{}
			path[depth] = next
			dfs(depth + 1)
			delete(used, next)
			extended = true
		}
		if !extended && depth > len(longest) {
			longest = append([]NodeHash(nil), path[:depth]...)
		}
	}

	for key := range data {
		for key2 := range data[key] {
			used[key] = struct{}{}
			used[key2] = struct{}{}
			path[0] = key
			path[1] = key2
			dfs(2)
			delete(used, key)
			delete(used, key2)
		}
	}
	return longest
}

func Part2(in io.Reader) string {
	items := parseInput(in)
	return StringifySequence(findLongest(items))
}
