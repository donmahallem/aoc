package day15

import (
	"bufio"
	"io"
)

type Lense struct {
	Id          uint64
	FocalLength uint64
}

/*
custom split function for bufio.Scanner to split by ','
*/
func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[0:i], nil
		}
	}
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func ParseInput(in io.Reader) [][]byte {
	scanner := bufio.NewScanner(in)
	scanner.Split(splitComma)
	groups := [][]byte{}
	for scanner.Scan() {
		token := append([]byte(nil), scanner.Bytes()...)
		groups = append(groups, token)
	}
	return groups
}

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
	scanner := bufio.NewScanner(in)
	scanner.Split(splitComma)

	memory := make([][]Lense, 256)

	for scanner.Scan() {
		data := scanner.Bytes()
		if data[len(data)-1] == '-' {
			hashedId, key := hashId(data[:len(data)-1])
			if box := memory[hashedId]; box != nil {
				for id := range box {
					if box[id].Id == key {
						box[id].Id = 0
					}
				}
			}
		} else if data[len(data)-2] == '=' {
			hashedId, key := hashId(data[:len(data)-2])
			value := uint64(data[len(data)-1] - '0')
			l1 := Lense{Id: key, FocalLength: value}
			if box := memory[hashedId]; box != nil {
				found := false
				for id := range box {
					if box[id].Id == key {
						box[id].FocalLength = value
						found = true
						break
					}
				}
				if !found {
					memory[hashedId] = append(box, l1)
				}
			} else {
				box := make([]Lense, 0, 4)
				box = append(box, l1)
				memory[hashedId] = box
			}
		}

	}
	accum := uint64(0)
	for idx, box := range memory {
		if box != nil {
			slotIndex := uint64(1)
			for _, lense := range box {
				if lense.Id != 0 {
					accum += uint64(idx+1) * slotIndex * uint64(lense.FocalLength)
					slotIndex++
				}
			}
		}
	}
	return accum
}
