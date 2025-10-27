package day20

import (
	"bufio"
	"bytes"
	"io"
)

// converts module ID from byte slice to int
func convertId(data []byte) int {
	id := 0
	for _, b := range data {
		if b < 'a' || b > 'z' {
			continue
		}
		id = id<<8 + int(b)
	}
	return id
}

type parsedInput struct {
	BroadcasterID int
	Broadcaster   broadcasterModule
	Modules       map[int]module
}

var arrowSeperator = []byte("->")

func ParseInput(r io.Reader) parsedInput {
	scanner := bufio.NewScanner(r)
	inp := parsedInput{
		Modules: make(map[int]module),
	}

	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}

		parts := bytes.SplitN(line, arrowSeperator, 2)
		if len(parts) != 2 {
			continue
		}

		leftPart := bytes.TrimSpace(parts[0])
		rightPart := bytes.TrimSpace(parts[1])

		targetIDs := bytes.Split(rightPart, []byte(","))
		targets := make([]int, 0, len(targetIDs))
		for _, t := range targetIDs {
			t = bytes.TrimSpace(t)
			if len(t) == 0 {
				continue
			}
			targets = append(targets, convertId(t))
		}

		if len(leftPart) == 0 {
			continue
		}

		switch leftPart[0] {
		case '%':
			moduleID := convertId(leftPart)
			inp.Modules[moduleID] = &flipFlopModule{targetIds: targets}
		case '&':
			moduleID := convertId(leftPart)
			inp.Modules[moduleID] = &conjunctionModule{targetIds: targets}
		default:
			if bytes.Equal(leftPart, []byte("broadcaster")) {
				inp.BroadcasterID = convertId(leftPart)
				inp.Broadcaster = broadcasterModule{targetIds: targets}
			}
		}
	}

	for srcID, mod := range inp.Modules {
		for _, targetID := range mod.TargetIds() {
			if conj, ok := inp.Modules[targetID].(*conjunctionModule); ok {
				conj.AddInput(srcID)
			}
		}
	}

	for _, targetID := range inp.Broadcaster.TargetIds() {
		if conj, ok := inp.Modules[targetID].(*conjunctionModule); ok {
			conj.AddInput(inp.BroadcasterID)
		}
	}

	for _, mod := range inp.Modules {
		mod.Reset()
	}

	return inp
}

type pulse struct {
	// pulse source
	from int
	// pulse target
	to int
	// pulse state
	high bool
}

// simulates a single press with optional observer callback
func simulatePress(inp *parsedInput, observer func(pulse)) (int, int) {
	lowCount, highCount := 1, 0

	queue := make([]pulse, 0, 32)
	for _, targetID := range inp.Broadcaster.TargetIds() {
		queue = append(queue, pulse{
			from: inp.BroadcasterID,
			to:   targetID,
			high: false,
		})
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if observer != nil {
			observer(current)
		}
		if current.high {
			highCount++
		} else {
			lowCount++
		}

		mod, ok := inp.Modules[current.to]
		if !ok {
			continue
		}

		emit, outgoingPulse := mod.Receive(current.from, current.high)
		if !emit {
			continue
		}

		for _, targetID := range mod.TargetIds() {
			queue = append(queue, pulse{
				from: current.to,
				to:   targetID,
				high: outgoingPulse,
			})
		}
	}

	return lowCount, highCount
}

func Part1(in io.Reader) int {
	inp := ParseInput(in)

	totalLow, totalHigh := 0, 0
	for range 1000 {
		low, high := simulatePress(&inp, nil)
		totalLow += low
		totalHigh += high
	}

	return totalLow * totalHigh
}
