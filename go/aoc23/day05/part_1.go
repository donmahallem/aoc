package day05

import (
	"bufio"
	"fmt"
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
	HumidityToLocation map[int]int
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
	var target *map[int]int
	data := make([]int, 0, 3)
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
			alma.SeedToSoil = make(map[int]int)
			target = &alma.SeedToSoil
		} else if strings.HasPrefix(line, SOIL_TO_FERTILIZER) {
			alma.SoilToFertilizer = make(map[int]int)
			target = &alma.SoilToFertilizer
		} else if strings.HasPrefix(line, FERTILIZER_TO_WATER) {
			alma.FertilizerToWater = make(map[int]int)
			target = &alma.FertilizerToWater
		} else if strings.HasPrefix(line, WATER_TO_LIGHT) {
			alma.WaterToLight = make(map[int]int)
			target = &alma.WaterToLight
		} else if strings.HasPrefix(line, LIGHT_TO_TEMPERATURE) {
			alma.LightToTemperature = make(map[int]int)
			target = &alma.LightToTemperature
		} else if strings.HasPrefix(line, TEMPERATURE_TO_HUMIDITY) {
			alma.TemperatureToHumidity = make(map[int]int)
			target = &alma.TemperatureToHumidity
		} else if strings.HasPrefix(line, HUMIDITY_TO_LOCATION) {
			alma.HumidityToLocation = make(map[int]int)
			target = &alma.HumidityToLocation
		} else {
			data = data[:0]
			ParseMap(&line, &data)
			TranslateMap(&data, target)
		}

	}
	return alma
}

func GetPosition(a *Almanac, start int) int {
	pos := start
	if pos1, ok := (*a).SeedToSoil[pos]; ok {
		pos = pos1
	}
	if pos1, ok := (*a).SoilToFertilizer[pos]; ok {
		pos = pos1
	}
	if pos1, ok := (*a).FertilizerToWater[pos]; ok {
		pos = pos1
	}
	if pos1, ok := (*a).WaterToLight[pos]; ok {
		pos = pos1
	}
	if pos1, ok := (*a).LightToTemperature[pos]; ok {
		pos = pos1
	}
	if pos1, ok := (*a).TemperatureToHumidity[pos]; ok {
		pos = pos1
	}
	if pos1, ok := (*a).HumidityToLocation[pos]; ok {
		pos = pos1
	}
	return pos
}

func Part1(in io.Reader) int {
	almanac := ParseAlmanac(in)
	fmt.Println("Parsed almanac")
	lowest := math.MaxInt
	for _, seed := range almanac.Seeds {
		lowest = aoc_utils.Min(lowest, GetPosition(&almanac, seed))
	}
	return lowest
}
