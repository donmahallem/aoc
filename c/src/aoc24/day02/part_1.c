#include "aoc24/day02/day02.h"

#include <string.h>

#define MAX_LINE_NUMS 256

aoc_error_t aoc24_day02_part1(FILE *in, aoc_result_t *out_result)
{
    char line[1024];
    int nums[MAX_LINE_NUMS];
    size_t n;
    int sum = 0;

    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    while (fgets(line, sizeof(line), in) != NULL)
    {
        parse_line_to_ints(line, nums, &n);
        if (n == 0)
            continue;
        if (check_line(nums, n))
            sum += 1;
    }

    if (ferror(in) != 0)
        return AOC_ERR_IO;

    AOC_RESULT_SET_I64(out_result, sum);
    return AOC_OK;
}
