#ifndef AOC24_DAY05_H
#define AOC24_DAY05_H
#include <stdio.h>
#include <stdlib.h>

#include "aoc_util/registry.h"

#define AOC24_DAY05_MAX_LINE_LENGTH 1024

#define MAX_PAGE_IDS_PER_LINE 128
struct page_ordering_rule
{
    int before;
    int after;
};
/**
 * Reads the ordering rules from file input until the first blank line
 */
static inline aoc_error_t parse_ordering_rules(FILE *in, struct page_ordering_rule **out_rules, int *out_rule_count)
{
    struct page_ordering_rule *page_ordering_rules = NULL;
    int rule_buffer_size = 0;
    int rule_buffer_capacity = 16;
    page_ordering_rules = malloc(sizeof(struct page_ordering_rule) * rule_buffer_capacity);
    if (page_ordering_rules == NULL)
        return AOC_ERR_NOMEM;

    char line[AOC24_DAY05_MAX_LINE_LENGTH];
    while (fgets(line, AOC24_DAY05_MAX_LINE_LENGTH, in) != NULL)
    {
        if (line[0] == '\n' || line[0] == '\r')
        {
            break;
        }

        if (rule_buffer_size >= rule_buffer_capacity)
        {
            rule_buffer_capacity *= 2;
            struct page_ordering_rule *new_buffer = realloc(page_ordering_rules, sizeof(struct page_ordering_rule) * rule_buffer_capacity);
            if (new_buffer == NULL)
            {
                free(page_ordering_rules);
                return AOC_ERR_NOMEM;
            }
            page_ordering_rules = new_buffer;
        }

        if (sscanf(line, "%d|%d", &page_ordering_rules[rule_buffer_size].before, &page_ordering_rules[rule_buffer_size].after) != 2)
        {
            free(page_ordering_rules);
            return AOC_ERR_PARSE;
        }
        rule_buffer_size++;
    }

    *out_rules = page_ordering_rules;
    *out_rule_count = rule_buffer_size;
    return AOC_OK;
}

/* Public API */
aoc_error_t aoc24_day05_part1(FILE *in, aoc_result_t *out_result);
aoc_error_t aoc24_day05_part2(FILE *in, aoc_result_t *out_result);

#endif
