package day19

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part1(r io.Reader) (int, error) {
	data := parseInput(r)

	var sum int
	for _, rating := range data.Ratings {
		current := "in"
		for steps := 0; current != actionAccept && current != actionReject; steps++ {
			// all are visited at most once
			if steps > len(data.Workflows) {
				return 0, aoc_utils.NewSolverErrorf("workflow %q is part of a cycle", current)
			}
			workflow, ok := data.Workflows[current]
			if !ok {
				current = actionReject
				break
			}
			next, ok := workflow.Next(rating)
			if !ok {
				current = actionReject
				break
			}
			current = next
		}

		if current == actionAccept {
			sum += rating.X + rating.M + rating.A + rating.S
		}
	}

	return sum, nil
}
