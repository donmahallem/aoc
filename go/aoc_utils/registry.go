package aoc_utils

type Registry struct {
	data map[PartSelector]interface{}
}

func NewRegistry() Registry {
	return Registry{data: make(map[PartSelector]interface{})}
}

func (v Registry) Register(year int, day int, part int, fn interface{}) {
	v.data[PartSelector{Year: year, Day: day, Part: part}] = fn
}

func (v Registry) GetPart(selection PartSelector) (interface{}, bool) {
	fn, ok := v.data[selection]
	if ok {
		return fn, true
	}
	return nil, false
}

type YearRegistryFunction = func(day int, fn ...interface{})

func (v Registry) CreateYearRegistry(year int) YearRegistryFunction {
	return func(day int, fn ...interface{}) {
		switch len(fn) {
		case 1:
			v.Register(year, day, 1, fn[0])
		case 2:
			v.Register(year, day, 2, fn[1])
		default:
			panic("Failed")
		}
	}
}

type PartSelector struct {
	Year int
	Day  int
	Part int
}
