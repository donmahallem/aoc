package day11

import (
	"io"
)

func Part2(in io.Reader) (uint64, error) {
	g, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	counter, err := newPathCounter(g)
	if err != nil {
		return 0, err
	}

	svrKey := encodeKey('s', 'v', 'r')
	fftKey := encodeKey('f', 'f', 't')
	dacKey := encodeKey('d', 'a', 'c')
	outKey := encodeKey('o', 'u', 't')

	// Path 2: svr -> fft -> dac -> out
	svrToFft := counter.count(svrKey, fftKey, dacKey, outKey)
	fftToDac := counter.count(fftKey, dacKey, outKey)
	dacToOut := counter.count(dacKey, outKey, fftKey)

	// Path 1: svr -> dac -> fft -> out
	svrToDac := counter.count(svrKey, dacKey, fftKey, outKey)
	dacToFft := counter.count(dacKey, fftKey, outKey)
	fftToOut := counter.count(fftKey, outKey, dacKey)

	return svrToDac*dacToFft*fftToOut + svrToFft*fftToDac*dacToOut, nil
}
