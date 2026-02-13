# AOC CPP

[![codecov](https://codecov.io/gh/donmahallem/aoc/graphs/badge.svg?token=meG36M3hel&component=module_cpp)](https://app.codecov.io/gh/donmahallem/aoc?components%5B0%5D=CPP)

## Tests

```
cmake -S . -B build -DCMAKE_BUILD_TYPE=Coverage
cmake --build build
ctest --test-dir build --output-on-failure
cmake --build build --target coverage
```
