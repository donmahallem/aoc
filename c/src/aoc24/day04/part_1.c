#include "aoc24/day04/day04.h"
#include <string.h>
#include <stdbool.h>

aoc_error_t aoc24_day04_part1(FILE *in, aoc_result_t *out_result)
{
    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    char lines[4][AOC24_DAY04_MAX_LINE_LENGTH];
    int line_len[4] = {0};

    int expectedLineLength = -1;
    int totalCount = 0;
    int lineIdx = 0;
    int curLineIdx;
    bool seen_content = false;

    while (fgets(lines[curLineIdx = lineIdx % 4], AOC24_DAY04_MAX_LINE_LENGTH, in) != NULL)
    {
        // trim newline/CR
        size_t len = strlen(lines[curLineIdx]);
        while (len > 0 && (lines[curLineIdx][len - 1] == '\n' || lines[curLineIdx][len - 1] == '\r'))
        {
            lines[curLineIdx][--len] = '\0';
        }
        line_len[curLineIdx] = (int)len;

        if (len == 0)
        {
            if (!seen_content)
            {
                // skip leading blank lines
                lineIdx++;
                continue;
            }
            break;
        }

        seen_content = true;

        if (expectedLineLength == -1)
            expectedLineLength = (int)len;
        else if ((int)len != expectedLineLength)
            return AOC_ERR_PARSE_UNEQUAL_LINES;

        /* scan trimmed characters only */
        for (int idx = 0; idx < (int)len; ++idx)
        {
            char c = lines[curLineIdx][idx];

            /* horizontal checks */
            if (c == 'X')
            {
                if (idx >= 3 && lines[curLineIdx][idx - 1] == 'M' && lines[curLineIdx][idx - 2] == 'A' && lines[curLineIdx][idx - 3] == 'S')
                    totalCount++;
            }
            else if (c == 'S')
            {
                if (idx >= 3 && lines[curLineIdx][idx - 1] == 'A' && lines[curLineIdx][idx - 2] == 'M' && lines[curLineIdx][idx - 3] == 'X')
                    totalCount++;
            }

            /* vertical/diagonal checks require 4 rows available */
            if (lineIdx >= 3)
            {
                /* vertical up */
                if (c == 'X')
                {
                    if (lines[(curLineIdx + 3) % 4][idx] == 'M' &&
                        lines[(curLineIdx + 2) % 4][idx] == 'A' &&
                        lines[(curLineIdx + 1) % 4][idx] == 'S')
                        totalCount++;
                    /* diagonal up-left */
                    if (idx >= 3 &&
                        lines[(curLineIdx + 3) % 4][idx - 1] == 'M' &&
                        lines[(curLineIdx + 2) % 4][idx - 2] == 'A' &&
                        lines[(curLineIdx + 1) % 4][idx - 3] == 'S')
                        totalCount++;
                    /* diagonal up-right */
                    if (idx + 3 < expectedLineLength &&
                        lines[(curLineIdx + 3) % 4][idx + 1] == 'M' &&
                        lines[(curLineIdx + 2) % 4][idx + 2] == 'A' &&
                        lines[(curLineIdx + 1) % 4][idx + 3] == 'S')
                        totalCount++;
                }
                else if (c == 'S')
                {
                    if (lines[(curLineIdx + 3) % 4][idx] == 'A' &&
                        lines[(curLineIdx + 2) % 4][idx] == 'M' &&
                        lines[(curLineIdx + 1) % 4][idx] == 'X')
                        totalCount++;
                    if (idx >= 3 && lines[(curLineIdx + 3) % 4][idx - 1] == 'A' &&
                        lines[(curLineIdx + 2) % 4][idx - 2] == 'M' &&
                        lines[(curLineIdx + 1) % 4][idx - 3] == 'X')
                        totalCount++;
                    if (idx + 3 < expectedLineLength &&
                        lines[(curLineIdx + 3) % 4][idx + 1] == 'A' &&
                        lines[(curLineIdx + 2) % 4][idx + 2] == 'M' &&
                        lines[(curLineIdx + 1) % 4][idx + 3] == 'X')
                        totalCount++;
                }
            }
        }

        lineIdx++;
    }

    if (ferror(in) != 0)
        return AOC_ERR_IO;

    AOC_RESULT_SET_I64(out_result, totalCount);
    return AOC_OK;
}
