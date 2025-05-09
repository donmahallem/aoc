cmake_minimum_required(VERSION 3.14)
if (CMAKE_VERSION VERSION_GREATER_EQUAL "3.24.0")
		cmake_policy(SET CMP0135 NEW)
endif()
project(AocSolver)
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/bin)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/bin)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/bin)

option(BUILD_TESTS "Build tests" ON)

add_subdirectory(./src)
#target_link_libraries(aoc_solver PUBLIC libaoc_util)

if(BUILD_TESTS)
  # Enable testing globally
  enable_testing()

  # Download and configure GoogleTest using FetchContent
  include(FetchContent)
  FetchContent_Declare(
    googletest
    URL https://github.com/google/googletest/archive/release-1.12.1.zip
  )
  # For Windows: Prevent overriding the parent's compiler/linker settings.
  set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
  FetchContent_MakeAvailable(googletest)

  # Add the test directory which builds your tests
  add_subdirectory(test)
endif()