# Build a single utility library that collects all sources under aoc_util/
file(GLOB_RECURSE AOC_UTIL_SRC CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/aoc_util/*.c)
add_library(aoc_util ${AOC_UTIL_SRC})
target_include_directories(aoc_util PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

# Find all aocXX directories
file(GLOB CHILDREN RELATIVE ${CMAKE_CURRENT_SOURCE_DIR} ${CMAKE_CURRENT_SOURCE_DIR}/aoc*)
set(YEAR_DIRS "")
foreach(child ${CHILDREN})
  if(IS_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/${child} AND child MATCHES "^aoc[0-9]+$")
    list(APPEND YEAR_DIRS ${child})
  endif()
endforeach()

set(ALL_YEAR_LIBS "")
set(ALL_REGISTER_INCLUDES "")
set(ALL_REGISTER_CALLS "")

foreach(YEAR_DIR ${YEAR_DIRS})
    string(SUBSTRING ${YEAR_DIR} 3 -1 YEAR_NUM)

    # Collect all sources for this year
    file(GLOB_RECURSE YEAR_SRC CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/${YEAR_DIR}/*.c)

    # Generate registration helper for this year
    file(GLOB DAY_HEADERS CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/${YEAR_DIR}/day*/day*.h)
    set(REGISTER_CALLS "")
    set(INCLUDES "")
    
    foreach(h ${DAY_HEADERS})
        file(RELATIVE_PATH rel ${CMAKE_CURRENT_SOURCE_DIR} ${h})
        string(REPLACE "\\" "/" rel ${rel})
        # Regex to match aocXX/dayYY/dayYY.h
        string(REGEX REPLACE "${YEAR_DIR}/day([0-9]+)/day[0-9]+\\.h" "\\1" DAY_NUM ${rel})
        
        if(DAY_NUM MATCHES "^[0-9]+$")
            string(APPEND INCLUDES "#include \"${rel}\"\n")
            string(APPEND REGISTER_CALLS "  err = aoc_registry_register_day(registry, ${YEAR_NUM}, ${DAY_NUM}, aoc${YEAR_NUM}_day${DAY_NUM}_part1, aoc${YEAR_NUM}_day${DAY_NUM}_part2);\n  if (err != AOC_OK) return err;\n")
        endif()
    endforeach()

    set(GEN_FILE ${CMAKE_CURRENT_BINARY_DIR}/generated_aoc${YEAR_NUM}_register.c)
    file(WRITE ${GEN_FILE} "#include \"aoc_util/registry.h\"\n")
    file(APPEND ${GEN_FILE} "${INCLUDES}\n")
    file(APPEND ${GEN_FILE} "aoc_error_t aoc${YEAR_NUM}_register_generated(aoc_registry_t* registry) {\n  aoc_error_t err;\n${REGISTER_CALLS}  return AOC_OK;\n}\n")

    list(APPEND YEAR_SRC ${GEN_FILE})

    # Create library for this year
    add_library(aoc${YEAR_NUM} ${YEAR_SRC})
    target_include_directories(aoc${YEAR_NUM} PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
    target_link_libraries(aoc${YEAR_NUM} PUBLIC aoc_util)
    
    list(APPEND ALL_YEAR_LIBS aoc${YEAR_NUM})
    
    # Prepare master registry lines
    string(APPEND ALL_REGISTER_INCLUDES "#include \"${YEAR_DIR}/aoc${YEAR_NUM}.h\"\n")
    string(APPEND ALL_REGISTER_CALLS "  err = aoc${YEAR_NUM}_register_parts(registry);\n  if (err != AOC_OK) return err;\n")

endforeach()

# Generate master registry
set(MASTER_GEN_FILE ${CMAKE_CURRENT_BINARY_DIR}/generated_all_years.c)
set(MASTER_GEN_HEADER ${CMAKE_CURRENT_BINARY_DIR}/generated_all_years.h)

file(WRITE ${MASTER_GEN_HEADER} "#ifndef GENERATED_ALL_YEARS_H\n#define GENERATED_ALL_YEARS_H\n\n#include \"aoc_util/registry.h\"\n\naoc_error_t aoc_register_all_years(aoc_registry_t* registry);\n\n#endif\n")

file(WRITE ${MASTER_GEN_FILE} "#include \"generated_all_years.h\"\n")
file(APPEND ${MASTER_GEN_FILE} "${ALL_REGISTER_INCLUDES}\n")
file(APPEND ${MASTER_GEN_FILE} "aoc_error_t aoc_register_all_years(aoc_registry_t* registry) {\n  aoc_error_t err;\n${ALL_REGISTER_CALLS}  return AOC_OK;\n}\n")

add_library(aoc_all_years ${MASTER_GEN_FILE})
target_include_directories(aoc_all_years PUBLIC ${CMAKE_CURRENT_BINARY_DIR})
target_link_libraries(aoc_all_years PUBLIC ${ALL_YEAR_LIBS} aoc_util)

add_executable(aoc_solver main.c)
target_link_libraries(aoc_solver PRIVATE aoc_all_years aoc_util)
