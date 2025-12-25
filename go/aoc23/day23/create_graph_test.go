package day23

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func Test_createGraph(t *testing.T) {
	t.Run("respect slope", func(t *testing.T) {
		t.Run("test sample", func(t *testing.T) {
			reader := strings.NewReader(testData)
			result, w, h, err := parseInput(reader, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			g, startIdx, endIdx := createGraph(result, w, h)
			if startIdx != 1 {
				t.Errorf(`Expected startIdx to be 1, got %d`, startIdx)
			}
			if endIdx != 527 {
				t.Errorf(`Expected endIdx to be 527, got %d`, endIdx)
			}
			if len(g) != 21 {
				t.Errorf(`Expected graph to have %d nodes, got %d`, 9, len(g))
			}
		})
	})

	t.Run("don't respect slope", func(t *testing.T) {
		t.Run("test sample", func(t *testing.T) {
			reader := strings.NewReader(testData)
			result, w, h, err := parseInput(reader, false)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			g, startIdx, endIdx := createGraph(result, w, h)
			if startIdx != 1 {
				t.Errorf(`Expected startIdx to be 1, got %d`, startIdx)
			}
			if endIdx != 527 {
				t.Errorf(`Expected endIdx to be 527, got %d`, endIdx)
			}

			if len(g) != 9 {
				t.Errorf(`Expected graph to have %d nodes, got %d`, 9, len(g))
			}

			t.Run("check that all nodes are bidirectional", func(t *testing.T) {
				for fromIdx, node := range g {
					for _, edge := range node.neighbors {
						toNode, found := g[edge.toIdx]
						if !found {
							t.Errorf(`Expected to find node %d`, edge.toIdx)
							continue
						}
						foundReverse := false
						for _, reverseEdge := range toNode.neighbors {
							if reverseEdge.toIdx == fromIdx {
								foundReverse = true
								break
							}
						}
						if !foundReverse {
							t.Errorf(`Expected to find reverse edge from %d to %d`, edge.toIdx, fromIdx)
						}
					}
				}
			})
		})
		t.Run("real sample", func(t *testing.T) {
			sampleData, testDataOk := test_utils.GetTestData(23, 23)
			if !testDataOk {
				t.SkipNow()
				return
			}
			reader := strings.NewReader(sampleData)
			result, w, h, err := parseInput(reader, false)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			g, startIdx, endIdx := createGraph(result, w, h)
			if startIdx != 1 {
				t.Errorf(`Expected startIdx to be 1, got %d`, startIdx)
			}
			if endIdx != 19879 {
				t.Errorf(`Expected endIdx to be 19879, got %d`, endIdx)
			}

			if len(g) != 36 {
				t.Errorf(`Expected graph to have %d nodes, got %d`, 9, len(g))
			}

			t.Run("check that all nodes are bidirectional", func(t *testing.T) {
				for fromIdx, node := range g {
					for _, edge := range node.neighbors {
						toNode, found := g[edge.toIdx]
						if !found {
							t.Errorf(`Expected to find node %d`, edge.toIdx)
							continue
						}
						foundReverse := false
						for _, reverseEdge := range toNode.neighbors {
							if reverseEdge.toIdx == fromIdx {
								foundReverse = true
								break
							}
						}
						if !foundReverse {
							t.Errorf(`Expected to find reverse edge from %d to %d`, edge.toIdx, fromIdx)
						}
					}
				}
			})
		})
	})
}
