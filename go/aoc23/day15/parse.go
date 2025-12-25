package day15

import (
	"bytes"
	"io"
)

type lens struct {
	Id          uint64
	FocalLength uint64
}

func parseInput(in io.Reader) [][]byte {
	raw, err := io.ReadAll(in)
	if err != nil {
		return nil
	}
	raw = trimTrailingNewlines(raw)
	if len(raw) == 0 {
		return nil
	}

	parts := bytes.Split(raw, []byte{','})
	tokens := make([][]byte, 0, len(parts))
	for _, p := range parts {
		p = bytes.TrimSpace(p)
		if len(p) == 0 {
			continue
		}
		tokens = append(tokens, p)
	}
	if len(tokens) == 0 {
		return nil
	}
	return tokens
}

// trimTrailingNewlines strips trailing CR/LF to simplify parsing.
func trimTrailingNewlines(raw []byte) []byte {
	for len(raw) > 0 {
		last := raw[len(raw)-1]
		if last != '\n' && last != '\r' {
			break
		}
		raw = raw[:len(raw)-1]
	}
	return raw
}

// hashId performs the AoC hash and builds a compact lens key in one pass.
// For robustness we limit the id used to the last 8 bytes to avoid shifting
// beyond 64 bits for long identifiers.
func hashId(id []byte) (uint32, uint64) {
	if len(id) == 0 {
		return 0, 0
	}
	const maxIDBytes = 8
	if len(id) > maxIDBytes {
		// Use the last maxIDBytes to keep suffix uniqueness.
		id = id[len(id)-maxIDBytes:]
	}
	var hash uint32 = 0
	var key uint64 = 0
	for _, b := range id {
		hash = (hash + uint32(b)) * 17 % 256
		key = (key << 8) | uint64(b)
	}
	return hash, key
}
