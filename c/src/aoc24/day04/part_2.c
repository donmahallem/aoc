#include "aoc24/day04/day04.h"
#include <string.h>
#include <stdbool.h>

/* Streaming, memory-friendly Part 2 parser with robust trimmed-line handling and explicit bounds checks. */

aoc_error_t aoc24_day04_part2(FILE *in, aoc_result_t *out_result)
{
    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    char lines[3][AOC24_DAY04_MAX_LINE_LENGTH];
    int line_len[3] = {0};

    int expectedLineLength = -1;
    int crossCount = 0;
    int lineIdx = 0;
    int curLineIdx;
    bool seen_content = false;

    while (fgets(lines[curLineIdx = lineIdx % 3], AOC24_DAY04_MAX_LINE_LENGTH, in) != NULL)
    {
        /* trim newline/CR */
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

        // when we have a full 3-line window available, check the middle row for crosses
        if (lineIdx >= 2)
        {
            const int mid = (curLineIdx + 2) % 3;
            const int top = (curLineIdx + 1) % 3;
            const int bot = curLineIdx;

            // only indices with neighbors on both sides can be centers
            for (int idx = 1; idx + 1 < expectedLineLength; ++idx)
            {
                if (lines[mid][idx] != 'A')
                    continue;

                // diagonal top-left / bot-right
                char tl = lines[top][idx - 1];
                char br = lines[bot][idx + 1];
                bool d1 = (tl == 'M' && br == 'S') || (tl == 'S' && br == 'M');

                // diagonal top-right / bot-left
                char tr = lines[top][idx + 1];
                char bl = lines[bot][idx - 1];
                bool d2 = (tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M');

                if (d1 && d2)
                    ++crossCount;
            }
        }

        lineIdx++;
    }

    if (ferror(in) != 0)
        return AOC_ERR_IO;

    AOC_RESULT_SET_I64(out_result, crossCount);
    return AOC_OK;
}
