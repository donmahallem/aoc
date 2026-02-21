// Unit tests for aoc24_day01_parse_input

#ifndef TEST_DATA_DIR
#define TEST_DATA_DIR ""
#endif

#include "unity.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "aoc_util/parse_int_list.h"

void setUp(void) {}
void tearDown(void) {}

/* helper for verifying both allocation and buffer-based parsers */
static void verify_parse(const char *input, const char *sep,
                         const int64_t *expected, size_t exp_n)
{
    /* allocation version */
    int64_t *arr = NULL;
    size_t n = 0;
    aoc_error_t err = aoc_util_parse_int_list(input, sep, &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(exp_n, n);
    for (size_t i = 0; i < exp_n; ++i)
        TEST_ASSERT_EQUAL_INT64(expected[i], arr[i]);
    free(arr);

    /* buffer version */
    int64_t buf[64];
    size_t n2 = 0;
    err = aoc_util_parse_int_list_to_buf(input, sep, buf, 64, &n2);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(exp_n, n2);
    for (size_t i = 0; i < exp_n; ++i)
        TEST_ASSERT_EQUAL_INT64(expected[i], buf[i]);
}

static void test_aoc_util_parse_int_list_comma_separated_simple(void)
{
    const int64_t exp[] = {1, 2, 3, 4, 5};
    verify_parse("1,2,3,4,5", ",", exp, 5);
}
static void test_aoc_util_parse_int_list_comma_separated_complex(void)
{
    const int64_t exp[] = {1, 2, 3, 4, 5};
    verify_parse(" 1, 2 , 3,4 ,5 ", ",", exp, 5);
}
static void test_aoc_util_parse_int_list_space_separated_simple(void)
{
    const int64_t exp[] = {1, 2, 3, 4, 5};
    verify_parse("1 2 3 4 5", " ", exp, 5);
}
static void test_aoc_util_parse_int_list_space_separated_complex(void)
{
    const int64_t exp[] = {1, 2, 3, 4, 5};
    verify_parse("  1  2 3 4 5  ", " ", exp, 5);
}

static void test_negative_values(void)
{
    const int64_t exp[] = {-1, -2};
    verify_parse("-1,-2", ",", exp, 2);
}
static void test_empty_input(void)
{
    const int64_t exp[] = {};
    verify_parse("", ",", exp, 0);
}

static void test_trailing_and_leading_separators(void)
{
    const int64_t exp[] = {1, 2};
    verify_parse(",1,2,", ",", exp, 2);
}

static void test_consecutive_separators(void)
{
    const int64_t exp[] = {1, 2};
    verify_parse("1,,2", ",", exp, 2);
}

static void test_mixed_space_tab(void)
{
    const int64_t exp[] = {1, 2, 3};
    verify_parse(" 1 \t 2 ,\t3 ", ",", exp, 3);
}

/* duplicate negative-values test removed; already covered above */

static void test_invalid_characters(void)
{
    int64_t *arr = (int64_t *)0xdeadbeef;
    size_t n = 0xdeadbeef;

    aoc_error_t err = aoc_util_parse_int_list("1,a,3", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_ERR_PARSE, err);
    TEST_ASSERT_EQUAL_PTR((int64_t *)0xdeadbeef, arr);
    TEST_ASSERT_EQUAL_UINT64(0xdeadbeef, n);
}

static void test_null_arguments(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list(NULL, ",", &arr, &n));
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list("1,2", NULL, &arr, &n));
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list("1,2", ",", NULL, &n));
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list("1,2", ",", &arr, NULL));

    /* same checks for the buffer variant */
    int64_t buf[4];
    n = 0;
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list_to_buf(NULL, ",", buf, 4, &n));
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list_to_buf("1,2", NULL, buf, 4, &n));
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list_to_buf("1,2", ",", NULL, 4, &n));
    TEST_ASSERT_EQUAL(AOC_ERR_NULL_ARG,
                      aoc_util_parse_int_list_to_buf("1,2", ",", buf, 4, NULL));
}

static void test_parse_to_buf_capacity(void)
{
    int64_t buf[2];
    size_t n = 0;
    aoc_error_t err = aoc_util_parse_int_list_to_buf("1,2,3", ",", buf, 2, &n);
    TEST_ASSERT_EQUAL(AOC_ERR_BUFFER_OVERFLOW, err);
    // n should remain untouched
    TEST_ASSERT_EQUAL_UINT64(0, n);
}

int main(void)
{
    UNITY_BEGIN();
    RUN_TEST(test_aoc_util_parse_int_list_comma_separated_simple);
    RUN_TEST(test_aoc_util_parse_int_list_comma_separated_complex);
    RUN_TEST(test_aoc_util_parse_int_list_space_separated_simple);
    RUN_TEST(test_aoc_util_parse_int_list_space_separated_complex);
    RUN_TEST(test_null_arguments);
    RUN_TEST(test_empty_input);
    RUN_TEST(test_trailing_and_leading_separators);
    RUN_TEST(test_consecutive_separators);
    RUN_TEST(test_negative_values);
    RUN_TEST(test_invalid_characters);
    RUN_TEST(test_mixed_space_tab);

    // parse to buffer tests
    RUN_TEST(test_parse_to_buf_capacity);

    return UNITY_END();
}