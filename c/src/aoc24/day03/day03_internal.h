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
        else if (c == ',' || c == ')')
        {
            return c;
        }
        else
        {
            return 0;
        }
    }
    return 0;
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

    int value = 0;
    int parsedNums = 0;
    int endChar;
    endChar = read_int(in, &value, &parsedNums);
    if (endChar != ',' || parsedNums < 1)
    {
        fseek(in, -1, SEEK_CUR);
        return 1;
    }
    *product = value;
    endChar = read_int(in, &value, &parsedNums);
    if (endChar != ')' || parsedNums < 1)
    {
        fseek(in, -1, SEEK_CUR);
        return 1;
    }
    *product *= value;
    return 0;
}
#endif