package queues_test

import (
	"fmt"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils/queues"
)

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	test := queues.NewFifoQueue[int](4)
	for i := range 6 {
		test.Add(i)
	}
	if result := test.Get(0); result != 2 {
		t.Errorf(`Expected index %d to be 2`, result)
	}
}

func BenchmarkAdd(b *testing.B) {
	testSizes := []uint{10, 100, 1000, 10000}
	for _, testSize := range testSizes {
		b.Run(fmt.Sprintf("test fifo size %d", testSize), func(b *testing.B) {
			test := queues.NewFifoQueue[int](testSize)
			for b.Loop() {

				for i := range 10000 * 2 {
					test.Add(i)
				}
			}
		})
	}

}
