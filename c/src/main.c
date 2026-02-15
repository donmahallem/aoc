#include <errno.h>
#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>

#include "generated_all_years.h"

static int parse_int_arg(const char *text, int *out_value)
{
    char *end_ptr = NULL;
    long value;

    if (text == NULL || out_value == NULL)
    {
        return 0;
    }

    errno = 0;
    value = strtol(text, &end_ptr, 10);
    if (errno != 0 || end_ptr == text || *end_ptr != '\0')
    {
        return 0;
    }

    *out_value = (int)value;
    return 1;
}

static int print_result(const aoc_result_t *result)
{
    size_t i;

    if (result == NULL)
    {
        return 0;
    }

    switch (result->kind)
    {
    case AOC_RESULT_I64:
        printf("Result: %" PRId64 "\n", result->value.i64);
        return 1;
    case AOC_RESULT_U64:
        printf("Result: %" PRIu64 "\n", result->value.u64);
        return 1;
    case AOC_RESULT_CSTR:
        printf("Result: %s\n", result->value.cstr == NULL ? "" : result->value.cstr);
        return 1;
    case AOC_RESULT_I64_ARRAY:
        printf("Result: [");
        for (i = 0; i < result->value.i64_array.length; ++i)
        {
            if (i > 0)
            {
                printf(", ");
            }
            printf("%" PRId64, result->value.i64_array.items[i]);
        }
        printf("]\n");
        return 1;
    case AOC_RESULT_I16_ARRAY:
        printf("Result: [");
        for (i = 0; i < result->value.i16_array.length; ++i)
        {
            if (i > 0)
            {
                printf(", ");
            }
            printf("%d", (int)result->value.i16_array.items[i]);
        }
        printf("]\n");
        return 1;
    case AOC_RESULT_CSTR_ARRAY:
        printf("Result: [");
        for (i = 0; i < result->value.cstr_array.length; ++i)
        {
            if (i > 0)
            {
                printf(", ");
            }
            printf("%s", result->value.cstr_array.items[i] == NULL ? "" : result->value.cstr_array.items[i]);
        }
        printf("]\n");
        return 1;
    default:
        return 0;
    }
}

int main(int argc, char **argv)
{
    aoc_registry_t registry;
    aoc_part_fn_t part_fn = NULL;
    aoc_result_t result;
    aoc_error_t err;
    int year;
    int day;
    int part;

    if (argc != 4)
    {
        fprintf(stderr, "Expected args <shortYear> <day> <part>\n");
        return 1;
    }

    if (!parse_int_arg(argv[1], &year) || !parse_int_arg(argv[2], &day) ||
        !parse_int_arg(argv[3], &part))
    {
        fprintf(stderr, "Invalid integer provided.\n");
        return 1;
    }

    aoc_registry_init(&registry);
    err = aoc_register_all_years(&registry);
    if (err != AOC_OK)
    {
        fprintf(stderr, "Failed to register parts: %s\n", aoc_error_to_string(err));
        return 1;
    }

    err = aoc_registry_get_part(&registry, year, day, part, &part_fn);
    if (err != AOC_OK)
    {
        fprintf(stderr, "Lookup failed: %s\n", aoc_error_to_string(err));
        return 1;
    }

    err = part_fn(stdin, &result);
    if (err != AOC_OK)
    {
        fprintf(stderr, "Execution failed: %s\n", aoc_error_to_string(err));
        return 1;
    }

    if (!print_result(&result))
    {
        fprintf(stderr, "Unsupported result type.\n");
        return 1;
    }

    return 0;
}
