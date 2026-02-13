cmake_minimum_required(VERSION 3.14)
add_library(aoc_util registry.cpp aoc_input_error.cpp)
target_include_directories(aoc_util PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/..)