module github.com/donmahallem/aoc/go

go 1.27

require (
	github.com/donmahallem/aggregate_benchmark v0.0.2
	github.com/donmahallem/aoc/test/testgen v0.0.0
)

require github.com/google/jsonschema-go v0.4.2 // indirect

replace github.com/donmahallem/aoc/test/testgen => ../test/testgen
