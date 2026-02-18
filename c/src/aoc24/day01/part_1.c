#include "aoc24/day01/day01.h"

#include <ctype.h>
#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include "aoc_util/int_util.h"
#include "aoc24/day01/day01_internal.h"

aoc_error_t aoc24_day01_part1(FILE *in, aoc_result_t *out_result)
{
    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    int64_t *left = NULL;
    int64_t *right = NULL;
    size_t n = 0;
    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    if (err != AOC_OK)
    {
        free(left);
        free(right);
        return err;
    }

    if (n == 0)
    {
        AOC_RESULT_SET_I64(out_result, 0);
        free(left);
        free(right);
        return AOC_OK;
    }

    qsort(left, n, sizeof(*left), aoc_util_int_cmp_i64);
    qsort(right, n, sizeof(*right), aoc_util_int_cmp_i64);

    int64_t total = 0;
    for (size_t i = 0; i < n; ++i)
    {
        int64_t diff = left[i] - right[i];
        if (diff < 0)
            diff = -diff;
        total += diff;
    }

    AOC_RESULT_SET_I64(out_result, total);
    free(left);
    free(right);
    return AOC_OK;
}
