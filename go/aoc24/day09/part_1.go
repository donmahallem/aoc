package day09

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func convertInput(inp []byte) ([]int16, error) {
	baseData := make([]int16, 0)
	for i := range inp {
		if inp[i] < '0' || inp[i] > '9' {
			return nil, aoc_utils.NewUnexpectedInputError(inp[i])
		}
		spec := inp[i] - '0'
		for range spec {
			if i%2 == 0 {
				baseData = append(baseData, int16((i / 2)))
			} else {
				baseData = append(baseData, -1)
			}
		}
	}
	return baseData, nil
}
func checkSum(data *[]int16) int {
	checkSum := 0
	for i := range len(*data) {
		if (*data)[i] >= 0 {
			checkSum += int((*data)[i]) * i
		}
	}
	return checkSum
}
func compactData(data *[]int16) {
	j := len(*data) - 1
	for i := range len(*data) {
		if (*data)[i] >= 0 {
			continue
		} else if i >= j {
			break
		}
		for ; j > i; j-- {
			if (*data)[j] >= 0 {
				(*data)[i] = (*data)[j]
				(*data)[j] = -1
				break
			}
		}
	}
}

func Part1(in io.Reader) (int, error) {
	data, _ := io.ReadAll(in)
	expandedData, err := convertInput(data)
	if err != nil {
		return 0, err
	}
	compactData(&expandedData)
	return checkSum(&expandedData), nil
}
