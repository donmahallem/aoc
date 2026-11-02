file(GLOB_RECURSE SRC_FILES CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/*.cpp)
add_library(aoc23 ${SRC_FILES})
target_include_directories(aoc23 PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/..)
target_link_libraries(aoc23 PUBLIC aoc_util)