#include "aoc24/day04/day04.h"

aoc_error_t aoc24_day04_part2(FILE *in, aoc_result_t *out_result)
{

    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    char lines[3][AOC24_DAY04_MAX_LINE_LENGTH];

    // check if line length varies
    int expectedLineLength = -1;
    int totalCount = 0;
    int lineIdx = 0;
    int crossCount = 0;
    int idx;
    int curLineIdx;
    while ((fgets(lines[curLineIdx = lineIdx % 3], AOC24_DAY04_MAX_LINE_LENGTH, in)) != NULL)
    {
        for (idx = 0; lines[curLineIdx][idx] != '\0'; idx++)
        {
            if (expectedLineLength >= 0 && idx > expectedLineLength)
            {
                return AOC_ERR_PARSE_UNEQUAL_LINES;
            }
            char c = lines[curLineIdx][idx];

            if (lineIdx >= 2)
            {

                int midLineIdx = (curLineIdx + 2) % 3;
                int topLineIdx = (curLineIdx + 1) % 3;
                int botLineIdx = curLineIdx;

                if (idx > 0 && idx < expectedLineLength - 1 && lines[midLineIdx][idx] == 'A')
                {
                    // Check diagonals
                    // M . S
                    // . A .
                    // M . S

                    // Diagonal 1 (Top-Left to Bottom-Right)
                    int d1_valid = 0;
                    char tl = lines[topLineIdx][idx - 1];
                    char br = lines[botLineIdx][idx + 1];
                    if ((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M'))
                    {
                        d1_valid = 1;
                    }

                    // Diagonal 2 (Top-Right to Bottom-Left)
                    int d2_valid = 0;
                    char tr = lines[topLineIdx][idx + 1];
                    char bl = lines[botLineIdx][idx - 1];
                    if ((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M'))
                    {
                        d2_valid = 1;
                    }

                    if (d1_valid && d2_valid)
                    {
                        totalCount++;
                    }
                }
            }
        }
        if (expectedLineLength == -1)
        {
            expectedLineLength = idx;
        }
        else if (idx < expectedLineLength - 1)
        {
            return AOC_ERR_PARSE_UNEQUAL_LINES;
        }
        lineIdx++;
    }
    if (ferror(in) != 0)
        return AOC_ERR_IO;

    AOC_RESULT_SET_I64(out_result, totalCount);
    return AOC_OK;
}
