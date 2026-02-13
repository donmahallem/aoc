# test/CMakeLists.txt

# Collect all test source files (or list them manually)
file(GLOB_RECURSE TEST_FILES ./**/*.cpp)

# Create the test executable.
add_executable(aoc_solver_test main_test.cpp ${TEST_FILES})

# Link the test executable with the main library and GoogleTest libraries.
target_link_libraries(aoc_solver_test PRIVATE gtest gtest_main aoc24)

# Define TEST_DATA_DIR so generated tests can locate data files.
# The extra escaped quotes ensure the value is a C string literal after macro expansion.
target_compile_definitions(aoc_solver_test PRIVATE TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\")

# Optionally add the src include directory for tests.
target_include_directories(aoc_solver_test PRIVATE ${CMAKE_SOURCE_DIR}/src)

# Enable GoogleTest to discover and register tests with CTest.
include(GoogleTest)
gtest_discover_tests(aoc_solver_test)
