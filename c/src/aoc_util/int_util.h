#ifndef AOC_UTIL_INT_UTIL_H
#define AOC_UTIL_INT_UTIL_H

#include <stddef.h>
#include <stdint.h>
#include <stdio.h>

/** Compares two int64_t values for example usage in qsort
 *
 * @param a Pointer to the first int64_t value
 * @param b Pointer to the second int64_t value
 * @return -1 if *a < *b, 1 if *a > *b, 0 if *a == *b
 */
int aoc_util_int_cmp_i64(const void *a, const void *b);

/**
 * Finds the index of the first element in the sorted array that is not less than the key.
 * The array must be sorted in ascending order.
 * @param arr Pointer to the first element of the array
 * @param n Number of elements in the array
 * @param key The value to search for
 * @return The index of the first element not less than the key
 */
size_t aoc_util_int_lower_bound_i64(const int64_t *arr, size_t n, int64_t key);

/**
 * Finds the index of the first element in the sorted array that is greater than the key.
 * The array must be sorted in ascending order.
 * @param arr Pointer to the first element of the array
 * @param n Number of elements in the array
 * @param key The value to search for
 * @return The index of the first element greater than the key
 */
size_t aoc_util_int_upper_bound_i64(const int64_t *arr, size_t n, int64_t key);

#endif /* AOC_UTIL_INT_UTIL_H */
