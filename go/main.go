package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/donmahallem/aoc/aoc23"
	"github.com/donmahallem/aoc/aoc24"
	"github.com/donmahallem/aoc/aoc_utils"
)

func main() {
	fmt.Print("AOC Solver\n")
	var parseError error
	partSelector := aoc_utils.PartSelector{}
	partSelector.Year, parseError = strconv.Atoi(os.Args[1])
	if parseError != nil {
		fmt.Println(parseError)
		return
	}
	partSelector.Day, parseError = strconv.Atoi(os.Args[2])
	if parseError != nil {
		fmt.Println(parseError)
		return
	}
	partSelector.Part, parseError = strconv.Atoi(os.Args[3])
	if parseError != nil {
		fmt.Println(parseError)
		return
	}
	fmt.Printf("Requested parsing %d-%d Part: %d\n", partSelector.Year, partSelector.Day, partSelector.Part)
	partRegistry := aoc_utils.NewRegistry()
	aoc23.RegisterParts(&partRegistry)
	aoc24.RegisterParts(&partRegistry)
	takeFun, ok := partRegistry.GetPart(partSelector)
	if !ok {
		fmt.Errorf("Could not find requested part %v", partSelector)
		return
	}
	function := reflect.ValueOf(takeFun)
	var startTime = time.Now()
	results := function.Call([]reflect.Value{reflect.ValueOf(os.Stdin)})
	var endTime = time.Now()
	// Checking the type of the first result for demonstration.
	res := results[0].Interface()
	switch v := res.(type) {
	case int:
		fmt.Println("Result is int:", v)
	case []int:
		fmt.Println("Result is []int:", v)
	default:
		fmt.Println("Unknown result type")
	}
	fmt.Printf("Took: %d\n", endTime.Sub(startTime).Microseconds())
}
