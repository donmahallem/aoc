#include "aoc24/day04/day04.h"

aoc_error_t aoc24_day04_part1(FILE *in, aoc_result_t *out_result)
{

    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    char lines[4][AOC24_DAY04_MAX_LINE_LENGTH];

    // check if line length varies
    int expectedLineLength = -1;
    int totalCount = 0;
    int lineIdx = 0;
    int idx;
    int curLineIdx;
    while (fgets(lines[curLineIdx = lineIdx % 4], AOC24_DAY04_MAX_LINE_LENGTH, in) != NULL)
    {
        for (idx = 0; lines[curLineIdx][idx] != '\0'; ++idx)
        {
            // ignore trailing empty lines. Just terminate on first empty line and ignore the rest of the file
            if (lines[curLineIdx][idx] == '\n' && idx == 0)
            {
                break;
            }
            if (expectedLineLength >= 0 && idx > expectedLineLength)
            {
                return AOC_ERR_PARSE_UNEQUAL_LINES;
            }
            char c = lines[curLineIdx][idx];
            // only heck horizontal matches for XMAS and SAMX
            if (c == 'X')
            {
                if (idx >= 3 && lines[curLineIdx][idx - 1] == 'M' && lines[curLineIdx][idx - 2] == 'A' && lines[curLineIdx][idx - 3] == 'S')
                {
                    totalCount++;
                }
            }
            else if (c == 'S')
            {
                if (idx >= 3 && lines[curLineIdx][idx - 1] == 'A' && lines[curLineIdx][idx - 2] == 'M' && lines[curLineIdx][idx - 3] == 'X')
                {
                    totalCount++;
                }
            }

            /* vertical/diagonal only when we already have 4 lines check Upwards */
            if (lineIdx >= 3)
            {
                /* Vertical: Check Up */
                if (c == 'X')
                {
                    // Vertical Up: S A M X (reading down) -> X M A S (reading up)
                    if (lines[(curLineIdx + 3) % 4][idx] == 'M' && lines[(curLineIdx + 2) % 4][idx] == 'A' && lines[(curLineIdx + 1) % 4][idx] == 'S')
                    {
                        totalCount++;
                    }
                    // Diagonal Up-Left
                    if (idx >= 3 && lines[(curLineIdx + 3) % 4][idx - 1] == 'M' && lines[(curLineIdx + 2) % 4][idx - 2] == 'A' && lines[(curLineIdx + 1) % 4][idx - 3] == 'S')
                    {
                        totalCount++;
                    }
                    // Diagonal Up-Right
                    if (lines[(curLineIdx + 3) % 4][idx + 1] == 'M' && lines[(curLineIdx + 2) % 4][idx + 2] == 'A' && lines[(curLineIdx + 1) % 4][idx + 3] == 'S')
                    {
                        totalCount++;
                    }
                }
                else if (c == 'S')
                {
                    // Vertical Up: X M A S (reading down) -> S A M X (reading up)
                    if (lines[(curLineIdx + 3) % 4][idx] == 'A' && lines[(curLineIdx + 2) % 4][idx] == 'M' && lines[(curLineIdx + 1) % 4][idx] == 'X')
                    {
                        totalCount++;
                    }
                    // Diagonal Up-Left
                    if (idx >= 3 && lines[(curLineIdx + 3) % 4][idx - 1] == 'A' && lines[(curLineIdx + 2) % 4][idx - 2] == 'M' && lines[(curLineIdx + 1) % 4][idx - 3] == 'X')
                    {
                        totalCount++;
                    }
                    // Diagonal Up-Right
                    if (lines[(curLineIdx + 3) % 4][idx + 1] == 'A' && lines[(curLineIdx + 2) % 4][idx + 2] == 'M' && lines[(curLineIdx + 1) % 4][idx + 3] == 'X')
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
