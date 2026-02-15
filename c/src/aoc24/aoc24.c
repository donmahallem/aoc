#include "aoc24/aoc24.h"

/*
 * Delegates registration to the generated registrar which is produced by CMake
 * by scanning the aoc24/day* directories. This makes adding new days as
 * simple as creating the day folder with part_1.c and part_2.c plus a header.
 */
extern aoc_error_t aoc24_register_generated(aoc_registry_t *registry);

aoc_error_t aoc24_register_parts(aoc_registry_t *registry)
{
    return aoc24_register_generated(registry);
}
