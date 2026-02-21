#ifndef AOC_UTIL_PARSE_INT_LIST_H
#define AOC_UTIL_PARSE_INT_LIST_H

#include "aoc_util/registry.h"

/**
 * Parses a single-character separated list of integers from the input string
 * 'str'.  Whitespace (spaces or tabs) is always ignored before and after each
 * number, regardless of what separator is specified; this yields consistent
 * behaviour and also allows the separator itself to be a whitespace character.
 *
 * The function allocates an array of int64_t which the caller must free.  On
 * error no allocation is returned and the output pointers are left unmodified.
 *
 * @param str The input string containing the list of integers.
 * @param seperator Pointer to a character containing the separator.  Only the
 *        first character is used; passing "\t" will treat tabs as the
 *        separator, for example.
 * @param out_arr Output parameter that will point to the allocated array of
 *        parsed integers.
 * @param out_size Output parameter that will hold the number of integers
 *        parsed.
 */
aoc_error_t aoc_util_parse_int_list(const char *str, const char *seperator, int64_t **out_arr, size_t *out_size);

#endif /* AOC_UTIL_PARSE_INT_LIST_H */