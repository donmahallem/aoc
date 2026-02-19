// Unit tests for aoc24_day01_parse_input

#ifndef TEST_DATA_DIR
#define TEST_DATA_DIR ""
#endif

#include "unity.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "aoc24/day01/day01_internal.h"

void setUp(void) {}
void tearDown(void) {}

static void test_day01_parse_input_valid(void)
{
    FILE *in = tmpfile();
    int64_t *left = NULL, *right = NULL;
    size_t n = 0;
    aoc_error_t err;

    TEST_ASSERT_NOT_NULL(in);
    fputs("3 4\n4 3\n2 5\n1 3\n3 9\n3 3\n", in);
    rewind(in);

    err = aoc24_day01_parse_input(in, &left, &right, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(6, n);
    TEST_ASSERT_EQUAL_INT64(3, left[0]);
    TEST_ASSERT_EQUAL_INT64(4, right[0]);
    TEST_ASSERT_EQUAL_INT64(4, left[1]);
    TEST_ASSERT_EQUAL_INT64(3, right[1]);
    TEST_ASSERT_EQUAL_INT64(2, left[2]);
    TEST_ASSERT_EQUAL_INT64(5, right[2]);
    TEST_ASSERT_EQUAL_INT64(1, left[3]);
    TEST_ASSERT_EQUAL_INT64(3, right[3]);
    TEST_ASSERT_EQUAL_INT64(3, left[4]);
    TEST_ASSERT_EQUAL_INT64(9, right[4]);
    TEST_ASSERT_EQUAL_INT64(3, left[5]);
    TEST_ASSERT_EQUAL_INT64(3, right[5]);

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
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(2, n);
    TEST_ASSERT_EQUAL_INT64(3, left[0]);
    TEST_ASSERT_EQUAL_INT64(4, right[0]);
    TEST_ASSERT_EQUAL_INT64(4, left[1]);
    TEST_ASSERT_EQUAL_INT64(3, right[1]);

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
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(0, n);
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
    TEST_ASSERT_EQUAL(AOC_ERR_PARSE, err);
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
    TEST_ASSERT_EQUAL(AOC_ERR_PARSE, err);
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
    TEST_ASSERT_EQUAL(AOC_ERR_PARSE, err);
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
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(1000, n);
    for (size_t i = 0; i < n; ++i)
    {
        TEST_ASSERT_EQUAL_INT64((int64_t)i, left[i]);
        TEST_ASSERT_EQUAL_INT64((int64_t)(1000 - i), right[i]);
    }

    free(left);
    free(right);
    fclose(in);
}

int main(void)
{
    UNITY_BEGIN();
    RUN_TEST(test_day01_parse_input_valid);
    RUN_TEST(test_day01_parse_input_skips_blanks);
    RUN_TEST(test_day01_parse_input_empty_only);
    RUN_TEST(test_day01_parse_input_invalid_characters);
    RUN_TEST(test_day01_parse_input_single_number_line);
    RUN_TEST(test_day01_parse_input_extra_trailing_junk);
    RUN_TEST(test_day01_parse_input_very_long_valid_input);
    return UNITY_END();
}