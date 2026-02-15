// Unit tests for aoc24_day01_parse_input

#ifndef TEST_DATA_DIR
#define TEST_DATA_DIR ""
#endif

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "aoc24/day01/day01.h"

static void test_day01_parse_input_valid(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;
    aoc_error_t err;

    assert(in != NULL);
    fputs("3 4\n4 3\n2 5\n1 3\n3 9\n3 3\n", in);
    rewind(in);

    err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_OK);
    assert(n == 6);
    assert(left[0] == 3 && right[0] == 4);
    assert(left[1] == 4 && right[1] == 3);
    assert(left[2] == 2 && right[2] == 5);
    assert(left[3] == 1 && right[3] == 3);
    assert(left[4] == 3 && right[4] == 9);
    assert(left[5] == 3 && right[5] == 3);

    free(left);
    free(right);
    fclose(in);
}

static void test_day01_parse_input_skips_blanks(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;

    fputs("   \n3 4\n\n4 3\n", in);
    rewind(in);

    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_OK);
    assert(n == 2);
    assert(left[0] == 3 && right[0] == 4);
    assert(left[1] == 4 && right[1] == 3);

    free(left);
    free(right);
    fclose(in);
}

static void test_day01_parse_input_empty_only(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;

    fputs("\n   \n", in);
    rewind(in);

    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_OK);
    assert(n == 0);
    /* parser still returns allocated buffers; free them */
    free(left);
    free(right);
    fclose(in);
}

static void test_day01_parse_input_invalid_characters(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;

    fputs("3 a\n4 3\n", in);
    rewind(in);

    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_ERR_PARSE);
    /* nothing allocated on error path (or freed by caller) */
    fclose(in);
}

static void test_day01_parse_input_single_number_line(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;

    fputs("3\n4 3\n", in);
    rewind(in);

    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_ERR_PARSE);
    fclose(in);
}

static void test_day01_parse_input_extra_trailing_junk(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;

    fputs("3 4 xyz\n4 3\n", in);
    rewind(in);

    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_ERR_PARSE);
    fclose(in);
}

static void test_day01_parse_input_very_long_valid_input(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;

    /* test dynamic resizing of output buffers */
    for (size_t i = 0; i < 1000; ++i)
    {
        fprintf(in, "%zu %zu\n", i, 1000 - i);
    }
    rewind(in);

    aoc_error_t err = aoc24_day01_parse_input(in, &left, &right, &n);
    assert(err == AOC_OK);
    assert(n == 1000);
    for (size_t i = 0; i < n; ++i)
    {
        assert(left[i] == (int64_t)i);
        assert(right[i] == (int64_t)(1000 - i));
    }

    free(left);
    free(right);
    fclose(in);
}

int main(void)
{
    test_day01_parse_input_valid();
    test_day01_parse_input_skips_blanks();
    test_day01_parse_input_empty_only();
    test_day01_parse_input_invalid_characters();
    test_day01_parse_input_single_number_line();
    test_day01_parse_input_extra_trailing_junk();
    test_day01_parse_input_very_long_valid_input();
    return 0;
}