#ifndef AOC24_DAY02_H
#define AOC24_DAY02_H

#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>

#include "aoc_util/registry.h"

/* Helpers implemented as static inline so both part_*.c files can use them. */
static inline void parse_line_to_ints(const char *line, int *out, size_t *out_len)
{
    *out_len = 0;
    if (line == NULL)
        return;
    char *end = NULL;
    const char *p = line;
    while (*p != '\0' && *p != '\n')
    {
        long v = strtol(p, &end, 10);
        if (end == p)
            break;
        out[(*out_len)++] = (int)v;
        p = end;
        while (*p == ' ' || *p == '\t')
            ++p;
    }
}

static inline int check_line(const int *line, size_t len)
{
    if (len < 2)
        return 0;
    int diff;
    int upwardDir = line[0] < line[1];
    for (size_t i = 1; i < len; ++i)
    {
        diff = line[i] - line[i - 1];
        if (diff == 0 || abs(diff) > 3)
            return 0;
        if ((diff < 0) == upwardDir)
            return 0;
    }
    return 1;
}

static inline int is_line_fixable(const int *line, size_t len)
{
    if (len < 2)
        return 0;
    int upwardDir;
    int lIndex, rIndex;
    int validLine;
    for (int skipIdx = -1; skipIdx < (int)len; ++skipIdx)
    {
        if (skipIdx == 0)
        {
            if (len < 3)
                continue;
            upwardDir = line[1] < line[2];
            lIndex = 1;
            rIndex = 2;
        }
        else if (skipIdx == 1)
        {
            if (len < 3)
                continue;
            upwardDir = line[0] < line[2];
            lIndex = 0;
            rIndex = 2;
        }
        else
        {
            upwardDir = line[0] < line[1];
            lIndex = 0;
            rIndex = 1;
        }
        validLine = 1;
        while (rIndex < (int)len)
        {
            int diff = line[lIndex] - line[rIndex];
            if (diff == 0 || abs(diff) > 3 || ((diff < 0) != upwardDir))
            {
                validLine = 0;
                break;
            }
            lIndex++;
            if (lIndex == skipIdx)
                lIndex++;
            rIndex++;
            if (rIndex == skipIdx)
                rIndex++;
        }
        if (validLine)
            return 1;
    }
    return 0;
}

/* Public API */
aoc_error_t aoc24_day02_part1(FILE *in, aoc_result_t *out_result);
aoc_error_t aoc24_day02_part2(FILE *in, aoc_result_t *out_result);

#endif
