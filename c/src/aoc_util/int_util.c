#include "aoc_util/int_util.h"

int aoc_util_int_cmp_i64(const void *a, const void *b)
{
    int64_t va = *(const int64_t *)a;
    int64_t vb = *(const int64_t *)b;
    if (va < vb)
        return -1;
    if (va > vb)
        return 1;
    return 0;
}

size_t aoc_util_int_lower_bound_i64(const int64_t *arr, size_t n, int64_t key)
{
    size_t lo = 0, hi = n;
    while (lo < hi)
    {
        size_t mid = lo + (hi - lo) / 2;
        if (arr[mid] < key)
            lo = mid + 1;
        else
            hi = mid;
    }
    return lo;
}

size_t aoc_util_int_upper_bound_i64(const int64_t *arr, size_t n, int64_t key)
{
    size_t lo = 0, hi = n;
    while (lo < hi)
    {
        size_t mid = lo + (hi - lo) / 2;
        if (arr[mid] <= key)
            lo = mid + 1;
        else
            hi = mid;
    }
    return lo;
}