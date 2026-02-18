#include "aoc24/day03/day03.h"

#include "aoc24/day03/day03_internal.h"

aoc_error_t aoc24_day03_part1(FILE *in, aoc_result_t *out_result)
{

    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    int sum = 0;
    int curChar;
    size_t linePos = 0;
    int value;
    int end, endChar, numsParsed;
    while ((curChar = fgetc(in)) != EOF)
    {
        if (curChar == 'm')
        {
            if (parse_mul(in, &value) == 0)
            {
                sum += value;
            }
        }
    }
    if (ferror(in) != 0)
        return AOC_ERR_IO;

    AOC_RESULT_SET_I64(out_result, sum);
    return AOC_OK;
}
