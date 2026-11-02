package generator

import (
	"go/format"
	"path"
	"strconv"
	"text/template"

	"github.com/donmahallem/aoc/test/testgen/model"
)

// Go renders go/aocYY/dayNN/generated_test.go.
var Go = Language{
	Template:   "generated_test.go.tmpl",
	OutputFile: func(d *DayInfo) string { return path.Join(d.YearDir, dayPkg(d), "generated_test.go") },
	Context:    func(d *DayInfo) any { return goContext{DayInfo: d, GoPackage: dayPkg(d)} },
	Funcs:      template.FuncMap{"backtick": func() string { return "`" }},
	Expected:   goExpected,
	Format:     format.Source,
}

// goContext is the template data for generated_test.go.tmpl.
type goContext struct {
	*DayInfo
	GoPackage string
}

func goExpected(r model.Result, t model.ResultType) string {
	itoa := func(v int64) string { return strconv.FormatInt(v, 10) }
	switch r.Kind {
	case model.ResultInt:
		return itoa(r.Int)
	case model.ResultString:
		return strconv.Quote(r.Str)
	case model.ResultInts:
		return string(t) + "{" + join(r.Ints, itoa) + "}"
	case model.ResultStrings:
		return "[]string{" + join(r.Strs, strconv.Quote) + "}"
	}
	panic("unsupported result kind " + strconv.Itoa(int(r.Kind)))
}
