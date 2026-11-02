package model

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"maps"
	"slices"
)

type TestCases map[string]TestCase

func (r *TestCases) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	if kind := dec.PeekKind(); kind != '{' {
		return fmt.Errorf("test cases must be an object keyed by name; got %v", kind)
	}
	if _, err := dec.ReadToken(); err != nil {
		return err
	}
	cases := make(TestCases)
	for dec.PeekKind() != '}' {
		nameTok, err := dec.ReadToken()
		if err != nil {
			return err
		}
		name := nameTok.String()
		if _, dup := cases[name]; dup {
			return fmt.Errorf("duplicate test name %q", name)
		}
		var item TestCase
		if err := json.UnmarshalDecode(dec, &item); err != nil {
			return err
		}
		cases[name] = item
	}
	if _, err := dec.ReadToken(); err != nil {
		return err
	}
	*r = cases
	return nil
}

// Encodes in name order because map iteration is random.
func (r *TestCases) MarshalJSONTo(enc *jsontext.Encoder) error {
	if *r == nil {
		return enc.WriteToken(jsontext.Null)
	}
	if err := enc.WriteToken(jsontext.BeginObject); err != nil {
		return err
	}
	for _, name := range slices.Sorted(maps.Keys(*r)) {
		if err := enc.WriteToken(jsontext.String(name)); err != nil {
			return err
		}
		if err := json.MarshalEncode(enc, (*r)[name]); err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndObject)
}
