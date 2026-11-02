package generator

import (
	"fmt"
	"math"
	"path"
	"strconv"

	"github.com/donmahallem/aoc/test/testgen/model"
)

// Cpp renders cpp/test/aocYY/dayNN/generated_test.cpp.
var Cpp = Language{
	Template:   "generated_test.cpp.tmpl",
	OutputFile: func(d *DayInfo) string { return path.Join(d.YearDir, dayPkg(d), "generated_test.cpp") },
	SourceFile: headerPath,
	Context: func(d *DayInfo) any {
		return cppContext{
			DayInfo:    d,
			HeaderPath: headerPath(d),
			Namespace:  d.YearDir + "::" + dayPkg(d),
			TestSuite:  fmt.Sprintf("Aoc%02dDay%02d", d.YearNumber, d.DayNumber),
		}
	},
	Expected: cppExpected,
}

// cppContext is the template data for generated_test.cpp.tmpl.
type cppContext struct {
	*DayInfo
	HeaderPath string
	Namespace  string
	TestSuite  string
}

func cppExpected(r model.Result, t model.ResultType) string {
	isInt16 := t.Elem() == model.ResultTypeInt16
	intLit := func(v int64) string {
		switch {
		case isInt16:
			return fmt.Sprintf("static_cast<int16_t>(%d)", v)
		case v > math.MaxInt32 || v < math.MinInt32:
			return fmt.Sprintf("%dLL", v)
		}
		return strconv.FormatInt(v, 10)
	}
	strLit := func(s string) string { return "std::string(" + strconv.Quote(s) + ")" }

	switch r.Kind {
	case model.ResultInt:
		return intLit(r.Int)
	case model.ResultString:
		return strLit(r.Str)
	case model.ResultInts:
		return "std::vector<" + cppElemType(t) + ">{" + join(r.Ints, intLit) + "}"
	case model.ResultStrings:
		return "std::vector<std::string>{" + join(r.Strs, strLit) + "}"
	}
	panic("unsupported result kind " + strconv.Itoa(int(r.Kind)))
}

func cppElemType(t model.ResultType) string {
	switch e := t.Elem(); e {
	case model.ResultTypeInt:
		return "int"
	case model.ResultTypeUint:
		return "unsigned"
	case model.ResultTypeString:
		return "std::string"
	default:
		return string(e) + "_t"
	}
}
