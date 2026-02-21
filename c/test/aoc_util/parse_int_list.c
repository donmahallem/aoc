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

static void test_aoc_util_parse_int_list_comma_separated_simple(void)
{
    const char *input = "1,2,3,4,5";
    int64_t *arr = NULL;
    size_t n = 0;
    aoc_error_t err;

    err = aoc_util_parse_int_list(input, ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(5, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    TEST_ASSERT_EQUAL_INT64(3, arr[2]);
    TEST_ASSERT_EQUAL_INT64(4, arr[3]);
    TEST_ASSERT_EQUAL_INT64(5, arr[4]);

    free(arr);
}
static void test_aoc_util_parse_int_list_comma_separated_complex(void)
{
    const char *input = " 1, 2 , 3,4 ,5 ";
    int64_t *arr = NULL;
    size_t n = 0;
    aoc_error_t err;

    err = aoc_util_parse_int_list(input, ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(5, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    TEST_ASSERT_EQUAL_INT64(3, arr[2]);
    TEST_ASSERT_EQUAL_INT64(4, arr[3]);
    TEST_ASSERT_EQUAL_INT64(5, arr[4]);

    free(arr);
}
static void test_aoc_util_parse_int_list_space_separated_simple(void)
{
    const char *input = "1 2 3 4 5";
    int64_t *arr = NULL;
    size_t n = 0;
    aoc_error_t err;

    err = aoc_util_parse_int_list(input, " ", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(5, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    TEST_ASSERT_EQUAL_INT64(3, arr[2]);
    TEST_ASSERT_EQUAL_INT64(4, arr[3]);
    TEST_ASSERT_EQUAL_INT64(5, arr[4]);

    free(arr);
}
static void test_aoc_util_parse_int_list_space_separated_complex(void)
{
    const char *input = "  1  2 3 4 5  ";
    int64_t *arr = NULL;
    size_t n = 0;
    aoc_error_t err;

    err = aoc_util_parse_int_list(input, " ", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(5, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    TEST_ASSERT_EQUAL_INT64(3, arr[2]);
    TEST_ASSERT_EQUAL_INT64(4, arr[3]);
    TEST_ASSERT_EQUAL_INT64(5, arr[4]);

    free(arr);
}

// additional edge-case tests
static void test_empty_input(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    aoc_error_t err = aoc_util_parse_int_list("", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(0, n);
    TEST_ASSERT_NOT_NULL(arr);
    free(arr);
}

static void test_trailing_and_leading_separators(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    aoc_error_t err = aoc_util_parse_int_list(",1,2,", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(2, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    free(arr);
}

static void test_consecutive_separators(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    aoc_error_t err = aoc_util_parse_int_list("1,,2", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(2, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    free(arr);
}

static void test_negative_values(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    aoc_error_t err = aoc_util_parse_int_list("-1,-2", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(2, n);
    TEST_ASSERT_EQUAL_INT64(-1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(-2, arr[1]);
    free(arr);
}

static void test_invalid_characters(void)
{
    int64_t *arr = (int64_t *)0xdeadbeef;
    size_t n = 0xdeadbeef;

    aoc_error_t err = aoc_util_parse_int_list("1,a,3", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_ERR_PARSE, err);
    TEST_ASSERT_EQUAL_PTR((int64_t *)0xdeadbeef, arr);
    TEST_ASSERT_EQUAL_UINT64(0xdeadbeef, n);
}

static void test_tab_separator(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    aoc_error_t err = aoc_util_parse_int_list("1\t2\t3", "\t", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(3, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    TEST_ASSERT_EQUAL_INT64(3, arr[2]);
    free(arr);
}

static void test_mixed_space_tab(void)
{
    int64_t *arr = NULL;
    size_t n = 0;

    /* should treat spaces and tabs identically, even around separators */
    aoc_error_t err = aoc_util_parse_int_list(" 1 \t 2 ,\t3 ", ",", &arr, &n);
    TEST_ASSERT_EQUAL(AOC_OK, err);
    TEST_ASSERT_EQUAL_UINT64(3, n);
    TEST_ASSERT_EQUAL_INT64(1, arr[0]);
    TEST_ASSERT_EQUAL_INT64(2, arr[1]);
    TEST_ASSERT_EQUAL_INT64(3, arr[2]);
    free(arr);
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
}

int main(void)
{
    UNITY_BEGIN();
    RUN_TEST(test_aoc_util_parse_int_list_comma_separated_simple);
    RUN_TEST(test_aoc_util_parse_int_list_comma_separated_complex);
    RUN_TEST(test_aoc_util_parse_int_list_space_separated_simple);
    RUN_TEST(test_aoc_util_parse_int_list_space_separated_complex);

    /* extra coverage */
    RUN_TEST(test_empty_input);
    RUN_TEST(test_trailing_and_leading_separators);
    RUN_TEST(test_consecutive_separators);
    RUN_TEST(test_negative_values);
    RUN_TEST(test_invalid_characters);
    RUN_TEST(test_tab_separator);
    RUN_TEST(test_mixed_space_tab);
    RUN_TEST(test_null_arguments);

    return UNITY_END();
}