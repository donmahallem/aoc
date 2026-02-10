package day11

import (
	"io"
)

type data struct {
	fftToDac uint64
	dacToFft uint64
	svrToFft uint64
	svrToDac uint64
	dacToOut uint64
	fftToOut uint64
}

func dfs(start uint64, target uint64, inputMap map[uint64][]uint64, cache map[uint64]uint64, ignoreA, ignoreB uint64) uint64 {
	if start == target {
		return 1
	}
	if start == ignoreA || start == ignoreB {
		return 0
	}
	if val, ok := cache[start]; ok {
		return val
	}
	currentValue := uint64(0)
	for _, next := range inputMap[start] {
		currentValue += dfs(next, target, inputMap, cache, ignoreA, ignoreB)
	}
	cache[start] = currentValue // Store the result in the cache
	return currentValue
}
func Part2(in io.Reader) (uint64, error) {
	inputMap, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	svrKey := uint64('s')<<16 + uint64('v')<<8 + uint64('r')
	fftKey := uint64('f')<<16 + uint64('f')<<8 + uint64('t')
	dacKey := uint64('d')<<16 + uint64('a')<<8 + uint64('c')
	outKey := uint64('o')<<16 + uint64('u')<<8 + uint64('t')

	// Path 2: svr -> fft -> dac -> out
	cache := make(map[uint64]uint64)
	connectionsSvrToFft := dfs(svrKey, fftKey, inputMap, cache, dacKey, outKey)
	clear(cache)
	connectionsFftToDac := dfs(fftKey, dacKey, inputMap, cache, outKey, 0)
	clear(cache)
	connectionsDacToOut := dfs(dacKey, outKey, inputMap, cache, fftKey, 0)
	clear(cache)

	// Path 1: svr -> dac -> fft -> out
	connectionsSvrToDac := dfs(svrKey, dacKey, inputMap, cache, fftKey, outKey)
	clear(cache)
	connectionsDacToFft := dfs(dacKey, fftKey, inputMap, cache, outKey, 0)
	clear(cache)
	connectionsFftToOut := dfs(fftKey, outKey, inputMap, cache, dacKey, 0)

	return connectionsSvrToDac*connectionsDacToFft*connectionsFftToOut +
		connectionsSvrToFft*connectionsFftToDac*connectionsDacToOut, nil
}
