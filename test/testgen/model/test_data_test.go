package model_test

import (
	"bytes"
	"encoding/json/jsontext"
	"reflect"
	"testing"

	"github.com/donmahallem/aoc/test/testgen/model"
)

func TestSymmetricMarshalling(t *testing.T) {
	original := model.TestData{
		2024: {
			1: model.TestCases{},
			2: model.TestCases{},
		},
	}

	var buf bytes.Buffer
	enc := jsontext.NewEncoder(&buf)
	if err := original.MarshalJSONTo(enc); err != nil {
		t.Fatalf("MarshalJSONTo failed: %v", err)
	}

	dec := jsontext.NewDecoder(&buf)
	var decoded model.TestData
	if err := decoded.UnmarshalJSONFrom(dec); err != nil {
		t.Fatalf("UnmarshalJSONFrom failed: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Errorf("Decoded data does not match original.\nOriginal: %+v\nDecoded: %+v", original, decoded)
	}
}
