add_executable(aoc_solver_bench bench_main.cpp)

include(FetchContent)

# Disable benchmark's own tests
set(BENCHMARK_ENABLE_TESTING OFF CACHE BOOL "" FORCE)
set(BENCHMARK_ENABLE_GTEST_TESTS OFF CACHE BOOL "" FORCE)
set(BENCHMARK_INSTALL_DOCS OFF CACHE BOOL "" FORCE)

FetchContent_Declare(
    googlebenchmark
    GIT_REPOSITORY https://github.com/google/benchmark.git
    GIT_TAG        v1.9.5
)
FetchContent_MakeAvailable(googlebenchmark)

# Disable nlohmann_json tests/install
set(JSON_BuildTests OFF CACHE BOOL "" FORCE)
set(JSON_Install OFF CACHE BOOL "" FORCE)

FetchContent_Declare(
    nlohmann_json
    GIT_REPOSITORY https://github.com/nlohmann/json.git
    GIT_TAG v3.12.0
)
FetchContent_MakeAvailable(nlohmann_json)

target_link_libraries(aoc_solver_bench PRIVATE
    benchmark::benchmark
    aoc24
    nlohmann_json::nlohmann_json
)

# Lock the output directory for the benchmark executable to bin/
target_compile_definitions(aoc_solver_bench PRIVATE
    TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\"
)