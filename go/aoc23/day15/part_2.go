package day15

import (
	"io"
)

type Lense struct {
	Id          uint64
	FocalLength uint64
}

func ParseInput(in io.Reader) [][]byte {
	raw, err := io.ReadAll(in)
	if err != nil {
		return nil
	}
	raw = trimTrailingNewlines(raw)
	if len(raw) == 0 {
		return nil
	}

	tokens := make([][]byte, 0, 64)
	start := 0
	for i := 0; i <= len(raw); i++ {
		if i < len(raw) && raw[i] != ',' {
			continue
		}
		if start < i {
			token := make([]byte, i-start)
			copy(token, raw[start:i])
			tokens = append(tokens, token)
		}
		start = i + 1
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
func hashId(id []byte) (uint32, uint64) {
	var hash uint32 = 0
	var key uint64 = 0
	for _, b := range id {
		hash = (hash + uint32(b)) * 17 % 256
		key = (key << 8) | uint64(b)
	}
	return hash, key
}

func Part2(in io.Reader) uint64 {
	raw, err := io.ReadAll(in)
	if err != nil {
		return 0
	}
	raw = trimTrailingNewlines(raw)
	if len(raw) == 0 {
		return 0
	}

	memory := make([][]Lense, 256)

	start := 0
	for i := 0; i <= len(raw); i++ {
		if i < len(raw) && raw[i] != ',' {
			continue
		}
		if start < i {
			token := raw[start:i]
			if token[len(token)-1] == '-' {
				// Removal path: locate the lens and splice it out of the box.
				hashedId, key := hashId(token[:len(token)-1])
				box := memory[hashedId]
				for idx, lense := range box {
					if lense.Id == key {
						memory[hashedId] = append(box[:idx], box[idx+1:]...)
						break
					}
				}
			} else {
				// Insertion/update path: parse focal length and update the box in place.
				eqIdx := -1
				for j := len(token) - 1; j >= 0; j-- {
					if token[j] == '=' {
						eqIdx = j
						break
					}
				}
				if eqIdx == -1 {
					start = i + 1
					continue
				}

				hashedId, key := hashId(token[:eqIdx])

				var focal uint64
				for j := eqIdx + 1; j < len(token); j++ {
					focal = focal*10 + uint64(token[j]-'0')
				}

				box := memory[hashedId]
				updated := false
				for idx := range box {
					if box[idx].Id == key {
						box[idx].FocalLength = focal
						updated = true
						break
					}
				}
				if !updated {
					memory[hashedId] = append(box, Lense{Id: key, FocalLength: focal})
				}
			}
		}
		start = i + 1
	}

	var accum uint64
	for idx, box := range memory {
		// Accumulate focusing power in box order.
		slot := uint64(1)
		for _, lense := range box {
			accum += uint64(idx+1) * slot * lense.FocalLength
			slot++
		}
	}
	return accum
}
