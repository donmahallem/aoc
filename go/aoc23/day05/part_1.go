package day05

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Almanac struct {
	SeedToSoil,
	SoilToFertilizer,
	FertilizerToWater,
	WaterToLight,
	LightToTemperature,
	TemperatureToHumidity,
	HumidityToLocation [][]int
	Seeds []int
}

const SEED_TO_SOIL string = "seed-to-soil"
const SOIL_TO_FERTILIZER string = "soil-to-fertilizer"
const FERTILIZER_TO_WATER string = "fertilizer-to-water"
const WATER_TO_LIGHT string = "water-to-light"
const LIGHT_TO_TEMPERATURE string = "light-to-temperature"
const TEMPERATURE_TO_HUMIDITY string = "temperature-to-humidity"
const HUMIDITY_TO_LOCATION string = "humidity-to-location"

func ParseMap(line *string, target *[]int) {
	a := strings.Split(*line, " ")
	for idx := range a {
		if len(a[idx]) == 0 {
			continue
		}
		val, _ := strconv.Atoi(a[idx])
		*target = append(*target, val)
	}
}

func TranslateMap(line *[]int, data *map[int]int) {
	for i := range (*line)[2] {
		(*data)[(*line)[1]+i] = (*line)[0] + i
	}
}

func ParseAlmanac(in io.Reader) Almanac {
	s := bufio.NewScanner(in)
	alma := Almanac{}
	var target *[][]int
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "seeds:") {
			alma.Seeds = make([]int, 0)
			data := line[6:]
			ParseMap(&data, &alma.Seeds)
			continue
		} else if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, SEED_TO_SOIL) {
			alma.SeedToSoil = make([][]int, 0)
			target = &alma.SeedToSoil
		} else if strings.HasPrefix(line, SOIL_TO_FERTILIZER) {
			alma.SoilToFertilizer = make([][]int, 0)
			target = &alma.SoilToFertilizer
		} else if strings.HasPrefix(line, FERTILIZER_TO_WATER) {
			alma.FertilizerToWater = make([][]int, 0)
			target = &alma.FertilizerToWater
		} else if strings.HasPrefix(line, WATER_TO_LIGHT) {
			alma.WaterToLight = make([][]int, 0)
			target = &alma.WaterToLight
		} else if strings.HasPrefix(line, LIGHT_TO_TEMPERATURE) {
			alma.LightToTemperature = make([][]int, 0)
			target = &alma.LightToTemperature
		} else if strings.HasPrefix(line, TEMPERATURE_TO_HUMIDITY) {
			alma.TemperatureToHumidity = make([][]int, 0)
			target = &alma.TemperatureToHumidity
		} else if strings.HasPrefix(line, HUMIDITY_TO_LOCATION) {
			alma.HumidityToLocation = make([][]int, 0)
			target = &alma.HumidityToLocation
		} else {
			data := make([]int, 0, 3)
			ParseMap(&line, &data)
			*target = append(*target, data)
			//TranslateMap(&data, target)
		}

	}
	return alma
}

func FindPos(k [][]int, pos int) int {
	for _, row := range k {
		if pos >= row[1] && pos < row[1]+row[2] {
			return pos + row[0] - row[1]
		}
	}
	return pos
}

func GetPosition(a Almanac, start int) int {
	pos := start
	pos = FindPos(a.SeedToSoil, pos)
	pos = FindPos(a.SoilToFertilizer, pos)
	pos = FindPos(a.FertilizerToWater, pos)
	pos = FindPos(a.WaterToLight, pos)
	pos = FindPos(a.LightToTemperature, pos)
	pos = FindPos(a.TemperatureToHumidity, pos)
	pos = FindPos(a.HumidityToLocation, pos)
	return pos
}

func Part1(in io.Reader) int {
	almanac := ParseAlmanac(in)
	lowest := math.MaxInt
	for _, seed := range almanac.Seeds {
		lowest = aoc_utils.Min(lowest, GetPosition(almanac, seed))
	}
	return lowest
}
