package day08

import (
	"bufio"
	"io"
	"sort"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type SourceData struct {
	mapping []int
	output  []int
}

func convertToBitMap(inp string) int {
	result := 0
	for _, char := range inp {
		result |= 1 << (char - 'a')
	}
	return result
}
func parseInput(inputData io.Reader) ([]SourceData, error) {
	var result []SourceData
	scanner := bufio.NewScanner(inputData)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		splitted := strings.Split(line, " | ")
		if len(splitted) != 2 {
			return nil, aoc_utils.NewParseErrorf("invalid input line: %s", line)
		}

		mappingStrs := strings.Fields(splitted[0])
		sort.Slice(mappingStrs, func(i, j int) bool {
			return len(mappingStrs[i]) < len(mappingStrs[j])
		})
		mapping := make([]int, len(mappingStrs))
		for i, s := range mappingStrs {
			mapping[i] = convertToBitMap(s)
		}

		outputStrs := strings.Fields(splitted[1])
		output := make([]int, len(outputStrs))
		for i, s := range outputStrs {
			output[i] = convertToBitMap(s)
		}

		result = append(result, SourceData{mapping: mapping, output: output})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func decodeMapping(mapping []int) map[int]int {
	one := mapping[0]
	seven := mapping[1]
	four := mapping[2]
	all5 := mapping[3] & mapping[4] & mapping[5]
	all6 := mapping[6] & mapping[7] & mapping[8]
	eight := mapping[9]

	segA := seven - one
	segF := one & all6
	segC := one - segF
	segD := (four - one) & all5
	segB := (four - one) - segD
	segG := all5 - segA - segD
	segE := eight - four - segA - segG

	return map[int]int{
		segA | segB | segC | segE | segF | segG:        0,
		segC | segF:                                    1,
		segA | segC | segD | segE | segG:               2,
		segA | segC | segD | segF | segG:               3,
		segB | segC | segD | segF:                      4,
		segA | segB | segD | segF | segG:               5,
		segA | segB | segD | segE | segF | segG:        6,
		segA | segC | segF:                             7,
		segA | segB | segC | segD | segE | segF | segG: 8,
		segA | segB | segC | segD | segF | segG:        9,
	}
}
