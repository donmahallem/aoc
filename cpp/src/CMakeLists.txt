cmake_minimum_required(VERSION 3.14)

add_subdirectory(./aoc_util)
add_subdirectory(./aoc24)

add_executable(aoc_solver main.cpp)
target_link_libraries(aoc_solver PRIVATE aoc_util aoc24)
target_include_directories(aoc_solver PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})