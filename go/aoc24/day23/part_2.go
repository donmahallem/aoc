package day23

import (
	"io"
	"slices"
)

type NodeHashList = []NodeHash

func StringifySequence(a []NodeHash) string {
	var out string = ""
	for _, item := range a {
		if len(out) > 0 {
			out += ","
		}
		out += string(*UnhashId(&item))
	}
	return out
}

func FindLongest(data *CombinationMap) []NodeHash {
	todo := make([]NodeHashList, 0)
	for key := range *data {
		for _, key2 := range (*data)[key] {
			todo = append(todo, []NodeHash{key, key2})
		}
	}
	longest := make([]NodeHash, 0)
	for len(todo) > 0 {
		current := todo[0]
		todo = todo[1:]
		//fmt.Printf("Current: %v\n", StringSequence(current))
		previousKey := current[len(current)-1]
		for _, key := range (*data)[previousKey] {
			//fmt.Printf("  Check key: %v\n", string(*UnhashId(&key)))
			valid := true
			for idx := range len(current) - 1 {
				if !slices.Contains((*data)[current[idx]], key) {
					valid = false
					break
				}
			}
			if valid {
				//fmt.Printf("    Append: %v\n", string(*UnhashId(&key)))
				tmpList := make([]NodeHash, len(current), len(current)+1)
				copy(tmpList, current)
				tmpList = append(tmpList, key)
				todo = append(todo, tmpList)
				if len(tmpList) > len(longest) {
					longest = tmpList
				}
			} /*			 else if len(longest) < len(current) {
				longest = append(longest, make([]NodeHash, len(current)-len(longest))...)
				copy(longest, current)
			}*/
		}
	}

	return longest
}

func Part2(in io.Reader) string {
	items := ParseInputMap(in)
	return StringifySequence(FindLongest(items))
}
