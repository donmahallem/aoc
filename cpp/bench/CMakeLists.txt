add_executable(aoc_solver_bench bench_main.cpp)

target_link_libraries(aoc_solver_bench PRIVATE
    benchmark::benchmark
    aoc24
    nlohmann_json::nlohmann_json
)

# Lock the output directory for the benchmark executable to bin/
target_compile_definitions(aoc_solver_bench PRIVATE
    TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\"
)
