cmake_minimum_required(VERSION 3.14)
add_executable(aoc_solver main.cpp)
add_subdirectory(./aoc_util)
target_link_libraries(aoc_solver PRIVATE aoc_util)