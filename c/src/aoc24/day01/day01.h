#ifndef AOC24_DAY01_H
#define AOC24_DAY01_H

#include <stdio.h>

#include "aoc_util/registry.h"
#include "aoc_util/int_util.h"

/**
 * Parses the input for day01
 *
 * @param in Input file stream to read from
 * @param out_left Output pointer for the left integer array (caller is responsible for freeing)
 * @param out_right Output pointer for the right integer array (caller is responsible for freeing)
 * @param out_n Output pointer for the number of integer pairs parsed
 * @return AOC_OK on success, or an appropriate error code on failure
 */
aoc_error_t aoc24_day01_parse_input(FILE *in, int64_t **out_left, int64_t **out_right, size_t *out_n);

aoc_error_t aoc24_day01_part1(FILE *in, aoc_result_t *out_result);
aoc_error_t aoc24_day01_part2(FILE *in, aoc_result_t *out_result);

#endif
