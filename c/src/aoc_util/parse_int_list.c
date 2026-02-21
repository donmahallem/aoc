#include <stdlib.h>
#include <ctype.h>

#include "aoc_util/parse_int_list.h"

/* parser treats spaces and tabs as generic whitespace; the supplied separator
   character is used in addition to, not instead of, whitespace. */
aoc_error_t aoc_util_parse_int_list(const char *str, const char *seperator, int64_t **out_arr, size_t *out_size)
{
    if (str == NULL || seperator == NULL || out_arr == NULL || out_size == NULL)
        return AOC_ERR_NULL_ARG;

    char sep = *seperator;

    size_t capacity = 16;
    size_t size = 0;
    int64_t *arr = malloc(sizeof(int64_t) * capacity);
    if (arr == NULL)
        return AOC_ERR_IO;

    const char *ptr = str;
    while (*ptr != '\0')
    {
        // always skip whitespace and separator characters
        if (*ptr == sep || isspace((unsigned char)*ptr))
        {
            ptr++;
            continue;
        }

        char *endptr;
        int64_t value = strtoll(ptr, &endptr, 10);
        if (ptr == endptr)
        {
            free(arr);
            return AOC_ERR_PARSE;
        }
        // consume any whitespace after the number
        while (isspace((unsigned char)*endptr))
            endptr++;

        // expand array if needed
        if (size >= capacity)
        {
            capacity *= 2;
            int64_t *new_arr = realloc(arr, sizeof(int64_t) * capacity);
            if (new_arr == NULL)
            {
                free(arr);
                return AOC_ERR_IO;
            }
            arr = new_arr;
        }

        arr[size++] = value;
        ptr = endptr;
    }

    *out_arr = arr;
    *out_size = size;
    return AOC_OK;
}

aoc_error_t aoc_util_parse_int_list_to_buf(const char *str, const char *seperator, int64_t *out_arr, size_t capacity, size_t *out_size)
{
    if (str == NULL || seperator == NULL || out_arr == NULL || out_size == NULL)
        return AOC_ERR_NULL_ARG;

    char sep = *seperator;
    size_t size = 0;
    const char *ptr = str;

    while (*ptr != '\0')
    {
        if (*ptr == sep || isspace((unsigned char)*ptr))
        {
            ptr++;
            continue;
        }

        char *endptr;
        int64_t value = strtoll(ptr, &endptr, 10);
        if (ptr == endptr)
        {
            return AOC_ERR_PARSE;
        }

        // consume any whitespace after the number
        while (isspace((unsigned char)*endptr))
            endptr++;

        if (size >= capacity)
        {
            return AOC_ERR_IO;
        }

        out_arr[size++] = value;
        ptr = endptr;
    }

    *out_size = size;
    return AOC_OK;
}
