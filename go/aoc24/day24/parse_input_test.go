package day24

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_hash(t *testing.T) {

}

func Test_parseInput(t *testing.T) {
	nodes := parseInput(strings.NewReader(testData))
	expected := []struct {
		id  nodeId
		val bool
	}{
		{id: HashIdA('x', '0', '1'), val: false},
		{id: HashIdA('x', '0', '2'), val: true},
		{id: HashIdA('x', '0', '3'), val: true},
		{id: HashIdA('x', '0', '4'), val: false},
	}
	for _, exp := range expected {
		if node, ok := nodes.inputNodes[exp.id]; !ok {
			t.Errorf("Expected node %v to be present", UnhashId(exp.id))
		} else if node != exp.val {
			t.Errorf("Expected node %v to have value %v, got %v", UnhashId(exp.id), exp.val, node)
		}
	}
}
