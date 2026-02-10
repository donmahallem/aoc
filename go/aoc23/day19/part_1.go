package day19

import (
	"io"
)

func Part1(r io.Reader) (int, error) {
	data := parseInput(r)

	var sum int
	for _, rating := range data.Ratings {
		current := "in"
		for current != actionAccept && current != actionReject {
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
