package day23

import (
	"bufio"
	"io"
)

type NodeHash = uint16
type CombinationMap = map[NodeHash]map[NodeHash]struct{}

func HashId(a []byte) NodeHash {
	return uint16(a[0])<<8 | uint16(a[1])
}

func UnhashId(id NodeHash) [2]byte {
	return [2]byte{byte(id >> 8), byte(id & 0xFF)}
}

func parseInput(in io.Reader) CombinationMap {
	points := make(CombinationMap, 128)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Bytes()
		key1 := HashId(line[0:2])
		key2 := HashId(line[3:5])
		if key1 > key2 {
			key2, key1 = key1, key2
		}
		if ca, ok := points[key1]; !ok {
			ca = make(map[NodeHash]struct{}, 6)
			ca[key2] = struct{}{}
			points[key1] = ca
		} else {
			ca[key2] = struct{}{}
		}
	}
	return points
}
