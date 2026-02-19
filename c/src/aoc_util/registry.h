#ifndef AOC_UTIL_REGISTRY_H
#define AOC_UTIL_REGISTRY_H

#include <stddef.h>
#include <stdint.h>
#include <stdio.h>

#define AOC_REGISTRY_MAX_ENTRIES 256

typedef enum aoc_error
{
    AOC_OK = 0,
    AOC_ERR_NULL_ARG,
    AOC_ERR_PARSE_UNEQUAL_LINES,
    AOC_ERR_PARSE,
    AOC_ERR_IO,
    AOC_ERR_REGISTRY_FULL,
    AOC_ERR_NOT_FOUND,
    AOC_ERR_INVALID_PART,
} aoc_error_t;

/**
 * Enumeration of possible result types that a part function can return.
 * The kind field in the aoc_result_t struct indicates which type of result is being returned, and the value union contains the actual result value.
 */
typedef enum aoc_result_kind
{
    AOC_RESULT_I64,
    AOC_RESULT_U64,
    AOC_RESULT_CSTR,
    AOC_RESULT_I64_ARRAY,
    AOC_RESULT_I16_ARRAY,
    AOC_RESULT_CSTR_ARRAY,
} aoc_result_kind_t;

/********************
 * Result value types
 ********************/

/**
 * Array of int64_t values. The items pointer is a pointer to an array of int64_t values, and the length field indicates the number of items in the array.
 */
typedef struct aoc_i64_array
{
    const int64_t *items;
    size_t length;
} aoc_i64_array_t;

/**
 * Array of int16_t values. The items pointer is a pointer to an array of int16_t values, and the length field indicates the number of items in the array.
 */
typedef struct aoc_i16_array
{
    const int16_t *items;
    size_t length;
} aoc_i16_array_t;

/**
 * Array of C strings. The items pointer is a pointer to an array of C string pointers, and the length field indicates the number of items in the array.
 */
typedef struct aoc_cstr_array
{
    const char *const *items;
    size_t length;
} aoc_cstr_array_t;

typedef struct aoc_result
{
    aoc_result_kind_t kind;
    /*
     * Union of possible result values. The actual type of the value is determined by the kind field.
     * The caller is responsible for ensuring that the value is interpreted correctly based on the kind.
     * For array types, the caller is responsible for ensuring that the items pointer remains valid for the lifetime of the result.
     */
    union
    {
        int64_t i64;
        uint64_t u64;
        const char *cstr;
        aoc_i64_array_t i64_array;
        aoc_i16_array_t i16_array;
        aoc_cstr_array_t cstr_array;
    } value;
} aoc_result_t;

/* Convenience macros for constructing results in day implementations. */
#define AOC_RESULT_SET_I64(res, v)       \
    do                                   \
    {                                    \
        (res)->kind = AOC_RESULT_I64;    \
        (res)->value.i64 = (int64_t)(v); \
    } while (0)
#define AOC_RESULT_SET_U64(res, v)        \
    do                                    \
    {                                     \
        (res)->kind = AOC_RESULT_U64;     \
        (res)->value.u64 = (uint64_t)(v); \
    } while (0)
#define AOC_RESULT_SET_CSTR(res, s)    \
    do                                 \
    {                                  \
        (res)->kind = AOC_RESULT_CSTR; \
        (res)->value.cstr = (s);       \
    } while (0)
#define AOC_RESULT_SET_I64_ARRAY(res, a, n)  \
    do                                       \
    {                                        \
        (res)->kind = AOC_RESULT_I64_ARRAY;  \
        (res)->value.i64_array.items = (a);  \
        (res)->value.i64_array.length = (n); \
    } while (0)
#define AOC_RESULT_SET_I16_ARRAY(res, a, n)  \
    do                                       \
    {                                        \
        (res)->kind = AOC_RESULT_I16_ARRAY;  \
        (res)->value.i16_array.items = (a);  \
        (res)->value.i16_array.length = (n); \
    } while (0)
#define AOC_RESULT_SET_CSTR_ARRAY(res, a, n)  \
    do                                        \
    {                                         \
        (res)->kind = AOC_RESULT_CSTR_ARRAY;  \
        (res)->value.cstr_array.items = (a);  \
        (res)->value.cstr_array.length = (n); \
    } while (0)

typedef aoc_error_t (*aoc_part_fn_t)(FILE *in, aoc_result_t *out_result);

typedef struct aoc_day_entry
{
    aoc_part_fn_t part1;
    aoc_part_fn_t part2;
} aoc_day_entry_t;

typedef struct aoc_registry_item
{
    int year;
    int day;
    aoc_day_entry_t entry;
} aoc_registry_item_t;

/**
 * Registry struct that allows registering part functions for each year and day.
 */
typedef struct aoc_registry
{
    /**
     * Array of registered entries. The registry does not take ownership of the function pointers.
     * The caller is responsible for ensuring that the function pointers remain valid for the lifetime of the registry.
     */
    aoc_registry_item_t items[AOC_REGISTRY_MAX_ENTRIES];
    /**
     * Number of registered entries in the registry. Must be less than or equal to AOC_REGISTRY_MAX_ENTRIES.
     */
    size_t count;
} aoc_registry_t;

/**
 * Initializes the registry by setting the count to 0.
 * The caller is responsible for ensuring that the registry pointer is valid.
 *
 * @param registry Pointer to the registry to initialize.
 */
void aoc_registry_init(aoc_registry_t *registry);

/**
 * Registers the part functions for the given year and day.
 * If the year and day already exist in the registry, updates the part functions.
 * If the registry is full, returns an appropriate error code.
 *
 * @param registry Pointer to the registry to update.
 * @param year The year of the challenge (e.g., 24 for 2024).
 * @param day The day of the challenge (1-25).
 * @param part1 Function pointer for part 1 of the challenge.
 * @param part2 Function pointer for part 2 of the challenge.
 * @return AOC_OK on success, or an appropriate error code on failure.
 */
aoc_error_t aoc_registry_register_day(
    aoc_registry_t *registry,
    int year,
    int day,
    aoc_part_fn_t part1,
    aoc_part_fn_t part2);

/**
 * Looks up the part function for the given year, day and part.
 * If found, stores the function in out_part and returns AOC_OK.
 * If not found, returns an appropriate error code.
 *
 * @param registry Pointer to the registry to search.
 * @param year The year of the challenge (e.g., 24 for 2024
 * @param day The day of the challenge (1-25).
 * @param part The part of the challenge (1 or 2).
 * @param out_part Output parameter to store the found function pointer.
 * @return AOC_OK on success, or an appropriate error code if not found or on error.
 */
aoc_error_t aoc_registry_get_part(
    const aoc_registry_t *registry,
    int year,
    int day,
    int part,
    aoc_part_fn_t *out_part);

/**
 * Converts an aoc_error_t value to a human-readable string.
 *
 * @param err The error code to convert.
 * @return A string representation of the error code.
 */
const char *aoc_error_to_string(aoc_error_t err);

#endif
