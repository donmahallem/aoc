package generator

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/donmahallem/aoc/test/testgen/model"
)

// C renders c/test/aocYY/dayNN_generated.c.
var C = Language{
	Template:   "generated_test.c.tmpl",
	OutputFile: func(d *DayInfo) string { return path.Join(d.YearDir, dayPkg(d)+"_generated.c") },
	SourceFile: headerPath,
	Context:    func(d *DayInfo) any { return cContext{DayInfo: d, HeaderPath: headerPath(d)} },
	Funcs:      template.FuncMap{"cEscapeString": cString},
	Expected:   cExpected,
}

// cContext is the template data for generated_test.c.tmpl.
type cContext struct {
	*DayInfo
	HeaderPath string
}

var cEscaper = strings.NewReplacer(`\`, `\\`, `"`, `\"`, "\n", `\n`, "\r", `\r`, "\t", `\t`)

// cString returns s as a quoted C string literal.
func cString(s string) string { return `"` + cEscaper.Replace(s) + `"` }

func cExpected(r model.Result, t model.ResultType) string {
	switch r.Kind {
	case model.ResultInt:
		if t == model.ResultTypeInt16 {
			return fmt.Sprintf("(int16_t)%d", r.Int)
		}
		return strconv.FormatInt(r.Int, 10)
	case model.ResultString:
		return cString(r.Str)
	case model.ResultInts:
		return "{" + join(r.Ints, func(v int64) string { return strconv.FormatInt(v, 10) }) + "}"
	case model.ResultStrings:
		return "{" + join(r.Strs, cString) + "}"
	}
	panic("unsupported result kind " + strconv.Itoa(int(r.Kind)))
}
