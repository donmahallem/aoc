# AOC C solver

Usage: `cat input.txt | ./aoc_solver {shortYear} {day} {part}`

## Build

```sh
cmake -S . -B ./build
cmake --build ./build
```

## Tests & Coverage

```sh
cmake -S . -B ./build -DCMAKE_BUILD_TYPE=Coverage
cmake --build ./build
ctest --test-dir ./build --output-on-failure
```

For coverage report generation `lcov` is required

```sh
cmake --build build --target coverage
```
