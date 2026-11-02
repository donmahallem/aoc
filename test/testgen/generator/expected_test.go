package generator

import (
	"testing"

	"github.com/donmahallem/aoc/test/testgen/model"
)

func TestGoExpected(t *testing.T) {
	tests := []struct {
		name string
		r    model.Result
		typ  model.ResultType
		want string
	}{
		{"int", model.Result{Kind: model.ResultInt, Int: 42}, model.ResultTypeInt, "42"},
		{"string", model.Result{Kind: model.ResultString, Str: `a"b`}, model.ResultTypeString, `"a\"b"`},
		{"ints", model.Result{Kind: model.ResultInts, Ints: []int64{1, 2}}, model.ResultTypeInt16s, "[]int16{1, 2}"},
		{"strings", model.Result{Kind: model.ResultStrings, Strs: []string{"a", "b"}}, model.ResultTypeStrings, `[]string{"a", "b"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goExpected(tt.r, tt.typ); got != tt.want {
				t.Errorf("goExpected() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCExpected(t *testing.T) {
	tests := []struct {
		name string
		r    model.Result
		typ  model.ResultType
		want string
	}{
		{"int16 scalar", model.Result{Kind: model.ResultInt, Int: 5}, model.ResultTypeInt16, "(int16_t)5"},
		{"int", model.Result{Kind: model.ResultInt, Int: 5}, model.ResultTypeInt, "5"},
		{"string", model.Result{Kind: model.ResultString, Str: "x"}, model.ResultTypeString, `"x"`},
		{"ints", model.Result{Kind: model.ResultInts, Ints: []int64{1, 2}}, model.ResultTypeInts, "{1, 2}"},
		{"strings", model.Result{Kind: model.ResultStrings, Strs: []string{"a", "b"}}, model.ResultTypeStrings, `{"a", "b"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cExpected(tt.r, tt.typ); got != tt.want {
				t.Errorf("cExpected() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCppExpected(t *testing.T) {
	tests := []struct {
		name string
		r    model.Result
		typ  model.ResultType
		want string
	}{
		{"int", model.Result{Kind: model.ResultInt, Int: 1}, model.ResultTypeInt, "1"},
		{"large int", model.Result{Kind: model.ResultInt, Int: 1 << 40}, model.ResultTypeInt64, "1099511627776LL"},
		{"int16 scalar", model.Result{Kind: model.ResultInt, Int: 5}, model.ResultTypeInt16, "static_cast<int16_t>(5)"},
		{"string", model.Result{Kind: model.ResultString, Str: "x"}, model.ResultTypeString, `std::string("x")`},
		{"ints", model.Result{Kind: model.ResultInts, Ints: []int64{1, 2}}, model.ResultTypeInts, "std::vector<int>{1, 2}"},
		{"strings", model.Result{Kind: model.ResultStrings, Strs: []string{"a"}}, model.ResultTypeStrings, `std::vector<std::string>{std::string("a")}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cppExpected(tt.r, tt.typ); got != tt.want {
				t.Errorf("cppExpected() = %q, want %q", got, tt.want)
			}
		})
	}
}
