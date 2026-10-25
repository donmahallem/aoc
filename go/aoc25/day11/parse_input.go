package day11

import (
	"bytes"
	"io"

	aoc_utils "github.com/donmahallem/aoc/go/aoc_utils"
)

var errNodeNameLength = aoc_utils.NewParseErrorf("node names must have 3 letters")

func encodeKey(a, b, c byte) uint32 {
	return uint32(a)<<16 | uint32(b)<<8 | uint32(c)
}

func parseInput(in io.Reader) (*graph, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, aoc_utils.NewParseError("reading input", err)
	}
	// Each line defines one source; sizing from the line count avoids map and slice growth.
	lines := bytes.Count(data, []byte{'\n'}) + 1
	ids := newNodeIndex(lines)
	node := ids.intern

	// pairs holds consecutive (from, to) node ids; capacity fits the puzzle input, append grows otherwise.
	pairs := make([]int32, 0, lines*6)
	for len(data) > 0 {
		var line []byte
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			line, data = data[:i], data[i+1:]
		} else {
			line, data = data, nil
		}
		if n := len(line); n > 0 && line[n-1] == '\r' {
			line = line[:n-1]
		}

		source := int32(-1)
		key := uint32(0)
		keyLen := 0
		for _, c := range line {
			switch {
			case c == ':':
				if keyLen != 3 {
					return nil, errNodeNameLength
				}
				source = node(key)
				key, keyLen = 0, 0
			case c >= 'a' && c <= 'z':
				key = key<<8 | uint32(c)
				keyLen++
			case c == ' ':
				if source >= 0 && keyLen > 0 {
					if keyLen != 3 {
						return nil, errNodeNameLength
					}
					pairs = append(pairs, source, node(key))
					key, keyLen = 0, 0
				}
			default:
				return nil, aoc_utils.NewUnexpectedInputError(c)
			}
		}
		if source < 0 {
			continue
		}
		if keyLen > 0 {
			if keyLen != 3 {
				return nil, errNodeNameLength
			}
			pairs = append(pairs, source, node(key))
		}
	}
	return buildGraph(ids, pairs), nil
}
