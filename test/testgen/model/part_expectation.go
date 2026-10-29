package model

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"fmt"
	"slices"
	"strings"
)

// ResultKind identifies which field of a Result holds the value.
type ResultKind int

const (
	ResultUnset ResultKind = iota
	ResultInt
	ResultInts
	ResultString
	ResultStrings
)

// Result is an expected answer: an int, []int, string, or []string.
type Result struct {
	Kind ResultKind
	Int  int64
	Ints []int64
	Str  string
	Strs []string
}

func (r *Result) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	switch dec.PeekKind() {
	case '0':
		r.Kind = ResultInt
		return json.UnmarshalDecode(dec, &r.Int)
	case '"':
		r.Kind = ResultString
		return json.UnmarshalDecode(dec, &r.Str)
	case '[':
		raw, err := dec.ReadValue()
		if err != nil {
			return err
		}
		if json.Unmarshal(raw, &r.Ints) == nil {
			r.Kind = ResultInts
			return nil
		}
		r.Ints = nil
		if json.Unmarshal(raw, &r.Strs) == nil {
			r.Kind = ResultStrings
			return nil
		}
		return errors.New("result array must contain only integers or only strings")
	default:
		return fmt.Errorf("result must be an integer, string, or array; got %v", dec.PeekKind())
	}
}

func (r Result) MarshalJSONTo(enc *jsontext.Encoder) error {
	switch r.Kind {
	case ResultInt:
		return json.MarshalEncode(enc, r.Int)
	case ResultInts:
		return json.MarshalEncode(enc, r.Ints)
	case ResultString:
		return json.MarshalEncode(enc, r.Str)
	case ResultStrings:
		return json.MarshalEncode(enc, r.Strs)
	default:
		return enc.WriteToken(jsontext.Null)
	}
}

// ResultType is the type hint from test_schema.json, e.g. "int" or "[]int16".
type ResultType string

const (
	ResultTypeInt     ResultType = "int"
	ResultTypeInt16   ResultType = "int16"
	ResultTypeInt32   ResultType = "int32"
	ResultTypeInt64   ResultType = "int64"
	ResultTypeUint    ResultType = "uint"
	ResultTypeUint16  ResultType = "uint16"
	ResultTypeUint32  ResultType = "uint32"
	ResultTypeUint64  ResultType = "uint64"
	ResultTypeString  ResultType = "string"
	ResultTypeInts    ResultType = "[]int"
	ResultTypeInt16s  ResultType = "[]int16"
	ResultTypeInt32s  ResultType = "[]int32"
	ResultTypeInt64s  ResultType = "[]int64"
	ResultTypeUints   ResultType = "[]uint"
	ResultTypeUint16s ResultType = "[]uint16"
	ResultTypeUint32s ResultType = "[]uint32"
	ResultTypeUint64s ResultType = "[]uint64"
	ResultTypeStrings ResultType = "[]string"
)

var resultTypes = []ResultType{
	ResultTypeInt, ResultTypeInt16, ResultTypeInt32, ResultTypeInt64,
	ResultTypeUint, ResultTypeUint16, ResultTypeUint32, ResultTypeUint64,
	ResultTypeString,
	ResultTypeInts, ResultTypeInt16s, ResultTypeInt32s, ResultTypeInt64s,
	ResultTypeUints, ResultTypeUint16s, ResultTypeUint32s, ResultTypeUint64s,
	ResultTypeStrings,
}

// IsArray reports whether the type is a slice type.
func (t ResultType) IsArray() bool { return strings.HasPrefix(string(t), "[]") }

// Elem returns the element type of an array type, or t itself for scalars.
func (t ResultType) Elem() ResultType { return ResultType(strings.TrimPrefix(string(t), "[]")) }

func (t *ResultType) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	var name string
	if err := json.UnmarshalDecode(dec, &name); err != nil {
		return err
	}
	// The schema allows "[]array" as an alias for a string array.
	if name == "[]array" {
		name = string(ResultTypeStrings)
	}
	if !slices.Contains(resultTypes, ResultType(name)) {
		return fmt.Errorf("unsupported type: %s", name)
	}
	*t = ResultType(name)
	return nil
}

// PartExpectation is the expected result of a single part.
type PartExpectation struct {
	Result        Result     `json:"result"`
	Type          ResultType `json:"type"`
	SkipLanguages []string   `json:"skip_languages,omitempty"`
}
