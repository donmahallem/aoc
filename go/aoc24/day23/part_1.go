package day23

import (
	"bufio"
	"io"
	"slices"
)

const PREFIX_T uint16 = uint16('t')

type NodeHash = uint16
type CombinationMap = map[NodeHash][]NodeHash

func HashId(a []byte) NodeHash {
	return uint16(a[0])<<8 | uint16(a[1])
}

func UnhashId(id *NodeHash) *[]byte {
	arr := make([]byte, 2)
	arr[0] = byte((*id) >> 8)
	arr[1] = byte((*id) & 255)
	return &arr
}

func ParseInputMap(in io.Reader) *CombinationMap {
	points := make(CombinationMap, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Bytes()
		key1 := HashId(line[0:2])
		key2 := HashId(line[3:5])
		if key1 > key2 {
			key2, key1 = key1, key2
		}
		if val, ok := points[key1]; ok {
			points[key1] = append(val, key2)
		} else {
			points[key1] = []uint16{key2}
		}
	}
	return &points
}

func StartsWithT(key *NodeHash) bool {
	return ((*key) >> 8) == PREFIX_T
}

func FindTriplets(data *CombinationMap) uint16 {
	results := uint16(0)
	var validEntry1, validEntry2 bool
	for key1 := range *data {
		dataKey1 := (*data)[key1]
		validEntry1 = StartsWithT(&key1)
		for _, key2 := range dataKey1 {
			validEntry2 = validEntry1 || StartsWithT(&key2)
			dataKey2 := (*data)[key2]
			for _, key3 := range dataKey2 {
				if validEntry2 || StartsWithT(&key3) {
					if slices.Contains(dataKey1, key3) {
						results++
					}
				}
			}
		}
	}
	return results
}

func Part1(in io.Reader) uint16 {
	items := ParseInputMap(in)
	return FindTriplets(items)
}
