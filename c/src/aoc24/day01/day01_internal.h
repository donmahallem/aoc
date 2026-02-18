#ifndef AOC24_DAY01_INTERNAL_H
#define AOC24_DAY01_INTERNAL_H

#include <stdio.h>
#include <ctype.h>
#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include "aoc_util/registry.h"

static inline aoc_error_t aoc24_day01_parse_input(FILE *in, int64_t **out_left, int64_t **out_right, size_t *out_n)
{
    char line[1024];
    size_t cap = 64;
    size_t n = 0;
    int64_t *left = malloc(cap * sizeof(*left));
    int64_t *right = malloc(cap * sizeof(*right));
    if (!left || !right)
    {
        free(left);
        free(right);
        return AOC_ERR_IO;
    }

    while (fgets(line, sizeof(line), in) != NULL)
    {
        /* skip leading whitespace */
        const char *p = line;
        while (isspace((unsigned char)*p))
            ++p;
        if (*p == '\0')
            continue; /* blank line */

        char *endptr = NULL;
        errno = 0;
        long long lv = strtoll(p, &endptr, 10);
        if (endptr == p)
            return AOC_ERR_PARSE; /* no number */
        if (lv < 0)
        {
            return AOC_ERR_PARSE; /* negative number */
        }

        /* move to second token */
        p = endptr;
        while (isspace((unsigned char)*p))
            ++p;
        if (*p == '\0')
            return AOC_ERR_PARSE; /* only one number on line */

        errno = 0;
        long long rv = strtoll(p, &endptr, 10);
        if (endptr == p)
            return AOC_ERR_PARSE; /* failed to parse second number */
        if (rv < 0)
            return AOC_ERR_PARSE; /* negative number */

        /* remainder after second number must be whitespace only */
        p = endptr;
        while (isspace((unsigned char)*p))
            ++p;
        if (*p != '\0')
            return AOC_ERR_PARSE; /* junk after two integers */

        if (n >= cap)
        {
            size_t newcap = cap * 2;

            int64_t *nl = realloc(left, newcap * sizeof(*nl));
            if (!nl)
            {
                free(left);
                free(right);
                return AOC_ERR_IO;
            }
            left = nl;

            int64_t *nr = realloc(right, newcap * sizeof(*nr));
            if (!nr)
            {
                free(left);
                free(right);
                return AOC_ERR_IO;
            }
            right = nr;

            cap = newcap;
        }

        left[n] = (int64_t)lv;
        right[n] = (int64_t)rv;
        ++n;
    }

    if (ferror(in))
    {
        free(left);
        free(right);
        return AOC_ERR_IO;
    }

    *out_left = left;
    *out_right = right;
    *out_n = n;
    return AOC_OK;
}
#endif /* AOC24_DAY01_INTERNAL_H */
