add_executable(aoc_solver_bench_c bench_main.cpp)

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

# bench_main.cpp is C++ wrapping C libraries, so set CXX standard on the target
set_target_properties(aoc_solver_bench_c PROPERTIES
    CXX_STANDARD 20
    CXX_STANDARD_REQUIRED ON
)

target_link_libraries(aoc_solver_bench_c PRIVATE
    benchmark::benchmark
    aoc_all_years
    nlohmann_json::nlohmann_json
)

target_compile_definitions(aoc_solver_bench_c PRIVATE
    TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\"
)
