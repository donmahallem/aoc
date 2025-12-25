package day05

import (
	"io"
	"math"
)

func findPos(k []almanacRange, pos int) int {
	for _, row := range k {
		if pos >= row.From.Min && pos < row.From.Max {
			return pos + row.To.Min - row.From.Min
		}
	}
	return pos
}

func getPosition(a almanac, start int) int {
	pos := start
	pos = findPos(a.SeedToSoil, pos)
	pos = findPos(a.SoilToFertilizer, pos)
	pos = findPos(a.FertilizerToWater, pos)
	pos = findPos(a.WaterToLight, pos)
	pos = findPos(a.LightToTemperature, pos)
	pos = findPos(a.TemperatureToHumidity, pos)
	pos = findPos(a.HumidityToLocation, pos)
	return pos
}

func Part1(in io.Reader) (int, error) {
	almanac := parseAlmanac(in)
	lowest := math.MaxInt
	for _, seed := range almanac.Seeds {
		lowest = min(lowest, getPosition(almanac, seed))
	}
	return lowest, nil
}
