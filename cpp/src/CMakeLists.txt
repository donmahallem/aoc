cmake_minimum_required(VERSION 3.14)

file(GLOB_RECURSE SRC_FILES ./**/*.cpp)
add_executable(aoc_solver main.cpp ${SRC_FILES})
add_subdirectory(./aoc_util)
add_subdirectory(./aoc24)
target_link_libraries(aoc_solver PUBLIC aoc_util aoc24)