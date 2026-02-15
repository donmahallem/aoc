cmake_minimum_required(VERSION 3.28)
project(AocSolverC C)

set(AOC_YEAR 24 CACHE STRING "Year of the Advent of Code event")

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)

option(BUILD_TESTS "Build tests" ON)

# Configure coverage flags if build type is Coverage.
# Must be set before add_subdirectory so all targets are instrumented.
if(CMAKE_BUILD_TYPE STREQUAL "Coverage")
  message(STATUS "Building in coverage mode")
  set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -O0 -g -fprofile-arcs -ftest-coverage -fno-inline")
  set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -g -fprofile-arcs -ftest-coverage")
endif()

add_subdirectory(src)

# Coverage target: prefer lcov + genhtml, fall back to gcovr, otherwise print a helpful message.
find_program(LCOV_EXEC lcov)
find_program(GENHTML_EXEC genhtml)

if(LCOV_EXEC AND GENHTML_EXEC)
  add_custom_target(coverage
    COMMAND ${CMAKE_CTEST_COMMAND} --test-dir ${CMAKE_BINARY_DIR} --output-on-failure
    COMMAND ${LCOV_EXEC} --capture --ignore-errors mismatch --directory ${CMAKE_BINARY_DIR} --output-file coverage_raw.info
    COMMAND ${LCOV_EXEC} --extract coverage_raw.info '${CMAKE_SOURCE_DIR}/src/*' --output-file coverage.info
    COMMAND ${GENHTML_EXEC} coverage.info --output-directory coverage_html
    WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
    COMMENT "Generating HTML coverage report with lcov/genhtml"
  )
else()
  add_custom_target(coverage
    COMMAND ${CMAKE_COMMAND} -E echo "Cannot generate coverage: 'lcov' was not found on PATH."
    COMMAND ${CMAKE_COMMAND} -E echo "Install 'lcov' + 'genhtml' to enable this target."
    COMMENT "coverage target (tooling missing)"
  )
endif()

if(BUILD_TESTS)
  enable_testing()
  add_subdirectory(test)
endif()
