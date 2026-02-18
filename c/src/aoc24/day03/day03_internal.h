#ifndef AOC24_DAY03_INTERNAL_H
#define AOC24_DAY03_INTERNAL_H

#include <stdio.h>

static inline int read_int(FILE *in, int *out, int *nums)
{
    *out = 0;
    *nums = 0;
    int c;
    while ((c = fgetc(in)) != EOF)
    {
        if (c >= '0' && c <= '9')
        {
            *out = (*out * 10) + (c - '0');
            (*nums)++;
        }
        else
        {
            return c;
        }
    }
    return EOF;
}

static inline int parse_mul(FILE *in, int *product)
{
    if (fgetc(in) != 'u')
    {
        fseek(in, -1, SEEK_CUR);
        return 1;
    }
    if (fgetc(in) != 'l')
    {
        fseek(in, -1, SEEK_CUR);
        return 1;
    }
    if (fgetc(in) != '(')
    {
        fseek(in, -1, SEEK_CUR);
        return 1;
    }

    int value1 = 0;
    int parsedCount1 = 0;
    int endChar;

    endChar = read_int(in, &value1, &parsedCount1);
    // Numbers can be 1-3 digits
    if (endChar != ',' || parsedCount1 < 1 || parsedCount1 > 3)
    {
        if (endChar != EOF)
        {
            fseek(in, -1, SEEK_CUR);
        }
        return 1;
    }

    int value2 = 0;
    int parsedCount2 = 0;

    endChar = read_int(in, &value2, &parsedCount2);
    if (endChar != ')' || parsedCount2 < 1 || parsedCount2 > 3)
    {
        if (endChar != EOF)
        {
            fseek(in, -1, SEEK_CUR);
        }
        return 1;
    }

    *product = value1 * value2;
    return 0;
}
#endif