package day05

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils"
)

// Interval half-open interval [Start, End)
type Interval struct {
	Start, End int
}

// Shift the interval
func (a *Interval) Shift(offset int) Interval {
	return Interval{Start: a.Start + offset, End: a.End + offset}
}

// Shift in place
func (a *Interval) ShiftI(offset int) {
	a.Start += offset
	a.End += offset
}

type AlmanacRange struct {
	From Interval
	To   Interval
}
type Almanac struct {
	SeedToSoil,
	SoilToFertilizer,
	FertilizerToWater,
	WaterToLight,
	LightToTemperature,
	TemperatureToHumidity,
	HumidityToLocation []AlmanacRange
	Seeds []int
}

const SEED_TO_SOIL string = "seed-to-soil"
const SOIL_TO_FERTILIZER string = "soil-to-fertilizer"
const FERTILIZER_TO_WATER string = "fertilizer-to-water"
const WATER_TO_LIGHT string = "water-to-light"
const LIGHT_TO_TEMPERATURE string = "light-to-temperature"
const TEMPERATURE_TO_HUMIDITY string = "temperature-to-humidity"
const HUMIDITY_TO_LOCATION string = "humidity-to-location"

func ParseMapping(line *string) AlmanacRange {
	a := strings.Split(*line, " ")
	val1, _ := strconv.Atoi(a[0])
	val2, _ := strconv.Atoi(a[1])
	val3, _ := strconv.Atoi(a[2])
	return AlmanacRange{From: Interval{Start: val2, End: val2 + val3}, To: Interval{Start: val1, End: val1 + val3}}
}

func ParseSeeds(line string) []int {
	a := strings.Split(line, " ")
	seeds := make([]int, len(a))
	for idx, num := range a {
		seeds[idx], _ = strconv.Atoi(num)
	}
	return seeds
}

func TranslateMap(line []int, data *map[int]int) {
	for i := range line[2] {
		(*data)[line[1]+i] = line[0] + i
	}
}

func ParseAlmanac(in io.Reader) Almanac {
	s := bufio.NewScanner(in)
	alma := Almanac{}
	var target *[]AlmanacRange
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "seeds:") {
			alma.Seeds = ParseSeeds(line[7:])
			continue
		} else if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, SEED_TO_SOIL) {
			alma.SeedToSoil = make([]AlmanacRange, 0)
			target = &alma.SeedToSoil
		} else if strings.HasPrefix(line, SOIL_TO_FERTILIZER) {
			alma.SoilToFertilizer = make([]AlmanacRange, 0)
			target = &alma.SoilToFertilizer
		} else if strings.HasPrefix(line, FERTILIZER_TO_WATER) {
			alma.FertilizerToWater = make([]AlmanacRange, 0)
			target = &alma.FertilizerToWater
		} else if strings.HasPrefix(line, WATER_TO_LIGHT) {
			alma.WaterToLight = make([]AlmanacRange, 0)
			target = &alma.WaterToLight
		} else if strings.HasPrefix(line, LIGHT_TO_TEMPERATURE) {
			alma.LightToTemperature = make([]AlmanacRange, 0)
			target = &alma.LightToTemperature
		} else if strings.HasPrefix(line, TEMPERATURE_TO_HUMIDITY) {
			alma.TemperatureToHumidity = make([]AlmanacRange, 0)
			target = &alma.TemperatureToHumidity
		} else if strings.HasPrefix(line, HUMIDITY_TO_LOCATION) {
			alma.HumidityToLocation = make([]AlmanacRange, 0)
			target = &alma.HumidityToLocation
		} else {
			*target = append(*target, ParseMapping(&line))
			//TranslateMap(&data, target)
		}

	}
	return alma
}

func FindPos(k []AlmanacRange, pos int) int {
	for _, row := range k {
		if pos >= row.From.Start && pos < row.From.End {
			return pos + row.To.Start - row.From.Start
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
