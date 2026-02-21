include(FetchContent)

# 1. Fetch Unity once before the loops
FetchContent_Declare(
    unity
    GIT_REPOSITORY https://github.com/ThrowTheSwitch/Unity.git
    GIT_TAG        v2.6.1
)
FetchContent_MakeAvailable(unity)

# Discover and build test targets for aoc_util

file(GLOB UTIL_TEST_SOURCES CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/aoc_util/*.c)
foreach(src IN LISTS UTIL_TEST_SOURCES)
    file(RELATIVE_PATH rel ${CMAKE_CURRENT_SOURCE_DIR} ${src})
    string(REPLACE "\\" "/" rel_unix ${rel})
    string(REPLACE "/" "_" rel_underscored ${rel_unix})
    string(REGEX REPLACE "[^A-Za-z0-9_]" "_" testname ${rel_underscored})
    set(target_name "c_test_${testname}")

    add_executable(${target_name} ${src})
    
    target_link_libraries(${target_name} PRIVATE 
        aoc_util 
        unity
    )

    target_compile_definitions(${target_name} PRIVATE TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\")
    add_test(NAME ${target_name} COMMAND ${target_name})
endforeach()

# Discover and build test targets for all aocXX directories

file(GLOB CHILDREN RELATIVE ${CMAKE_CURRENT_SOURCE_DIR} ${CMAKE_CURRENT_SOURCE_DIR}/aoc*)

foreach(child ${CHILDREN})
    if(IS_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/${child} AND child MATCHES "^aoc([0-9]+)$")
        set(YEAR_DIR ${child})
        set(YEAR_NUM ${CMAKE_MATCH_1})
        
        # Discover year-level C test sources
        file(GLOB TEST_SOURCES CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/${YEAR_DIR}/*.c)
        
        # Also pick up generator-produced files in per-day subfolders (generated_test.c)
        file(GLOB_RECURSE GENERATED_SOURCES CONFIGURE_DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/${YEAR_DIR}/*/generated_test.c)
        
        list(APPEND TEST_SOURCES ${GENERATED_SOURCES})

        foreach(src IN LISTS TEST_SOURCES)
            file(RELATIVE_PATH rel ${CMAKE_CURRENT_SOURCE_DIR} ${src})
            string(REPLACE "\\" "/" rel_unix ${rel})
            string(REPLACE "/" "_" rel_underscored ${rel_unix})
            string(REGEX REPLACE "[^A-Za-z0-9_]" "_" testname ${rel_underscored})
            set(target_name "c_test_${testname}")

            add_executable(${target_name} ${src})
            
            # 2. Add 'unity' to your library links
            target_link_libraries(${target_name} PRIVATE 
                aoc${YEAR_NUM} 
                aoc_util 
                unity
            )

            target_compile_definitions(${target_name} PRIVATE TEST_DATA_DIR=\"${CMAKE_SOURCE_DIR}/../test\")
            add_test(NAME ${target_name} COMMAND ${target_name})
        endforeach()
    endif()
endforeach()