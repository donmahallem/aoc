package day09

import (
	"fmt"
	"io"
)

func ConvertInput(inp *[]byte) []int16 {
	baseData := make([]int16, 0)
	for i := range len(*inp) {
		spec := (*inp)[i] - '0'
		for range spec {
			if i%2 == 0 {
				baseData = append(baseData, int16((i / 2)))
			} else {
				baseData = append(baseData, -1)
			}
		}
	}
	return baseData
}
func CheckSum(data *[]int16) int {
	checkSum := 0
	for i := range len(*data) {
		if (*data)[i] >= 0 {
			checkSum += int((*data)[i]) * i
		}
	}
	return checkSum
}
func CompactData(data *[]int16) {
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

func Part1(in io.Reader) {
	data, _ := io.ReadAll(in)
	expandedData := ConvertInput(&data)
	CompactData(&expandedData)
	fmt.Printf("%d\n", CheckSum(&expandedData))
}
