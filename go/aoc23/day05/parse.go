package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils/math"
)

type almanacInterval = math.Interval[int]
type almanacRange struct {
	From almanacInterval
	To   almanacInterval
}
type almanac struct {
	SeedToSoil,
	SoilToFertilizer,
	FertilizerToWater,
	WaterToLight,
	LightToTemperature,
	TemperatureToHumidity,
	HumidityToLocation []almanacRange
	Seeds []int
}

const seedToSoil string = "seed-to-soil"
const soilToFertilizer string = "soil-to-fertilizer"
const fertilizerToWater string = "fertilizer-to-water"
const waterToLight string = "water-to-light"
const lightToTemperature string = "light-to-temperature"
const temperatureToHumidity string = "temperature-to-humidity"
const humidityToLocation string = "humidity-to-location"

func parseMapping(line *string) almanacRange {
	// defensive parsing: a fuzzed line may not contain three space-separated
	// numeric tokens. In that case return a zero-valued AlmanacRange instead
	// of panicking.
	a := strings.Split(*line, " ")
	if len(a) < 3 {
		return almanacRange{}
	}
	val1, _ := strconv.Atoi(a[0])
	val2, _ := strconv.Atoi(a[1])
	val3, _ := strconv.Atoi(a[2])
	return almanacRange{From: almanacInterval{Min: val2, Max: val2 + val3}, To: almanacInterval{Min: val1, Max: val1 + val3}}
}

func parseSeeds(line string) []int {
	a := strings.Split(line, " ")
	seeds := make([]int, len(a))
	for idx, num := range a {
		seeds[idx], _ = strconv.Atoi(num)
	}
	return seeds
}

func translateMap(line []int, data *map[int]int) {
	for i := range line[2] {
		(*data)[line[1]+i] = line[0] + i
	}
}

func parseAlmanac(in io.Reader) almanac {
	s := bufio.NewScanner(in)
	alma := almanac{}
	var target *[]almanacRange
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "seeds:") {
			alma.Seeds = parseSeeds(line[7:])
			continue
		} else if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, seedToSoil) {
			alma.SeedToSoil = make([]almanacRange, 0)
			target = &alma.SeedToSoil
		} else if strings.HasPrefix(line, soilToFertilizer) {
			alma.SoilToFertilizer = make([]almanacRange, 0)
			target = &alma.SoilToFertilizer
		} else if strings.HasPrefix(line, fertilizerToWater) {
			alma.FertilizerToWater = make([]almanacRange, 0)
			target = &alma.FertilizerToWater
		} else if strings.HasPrefix(line, waterToLight) {
			alma.WaterToLight = make([]almanacRange, 0)
			target = &alma.WaterToLight
		} else if strings.HasPrefix(line, lightToTemperature) {
			alma.LightToTemperature = make([]almanacRange, 0)
			target = &alma.LightToTemperature
		} else if strings.HasPrefix(line, temperatureToHumidity) {
			alma.TemperatureToHumidity = make([]almanacRange, 0)
			target = &alma.TemperatureToHumidity
		} else if strings.HasPrefix(line, humidityToLocation) {
			alma.HumidityToLocation = make([]almanacRange, 0)
			target = &alma.HumidityToLocation
		} else {
			if target == nil {
				// malformed input: mapping line without a preceding section header
				continue
			}
			*target = append(*target, parseMapping(&line))
			//TranslateMap(&data, target)
		}

	}
	return alma
}
