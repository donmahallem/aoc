# test/CMakeLists.txt

# Collect all test source files.
file(GLOB_RECURSE TEST_FILES CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/*.cpp)

# Create the test executable. gtest_main provides main().
add_executable(aoc_solver_test ${TEST_FILES})

# Link the test executable with the main library and GoogleTest libraries.
target_link_libraries(aoc_solver_test PRIVATE gtest gtest_main aoc24)

# Define TEST_DATA_DIR so generated tests can locate data files.
# The extra escaped quotes ensure the value is a C string literal after macro expansion.
target_compile_definitions(aoc_solver_test PRIVATE TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\")

# Enable GoogleTest to discover and register tests with CTest.
include(GoogleTest)
gtest_discover_tests(aoc_solver_test)
