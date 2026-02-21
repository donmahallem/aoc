#include "aoc24/day05/day05.h"

#include <ctype.h>
#include "aoc_util/registry.h"
#include "aoc_util/parse_int_list.h"

aoc_error_t aoc24_day05_part1(FILE *in, aoc_result_t *out_result)
{
    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;
    char line[AOC24_DAY05_MAX_LINE_LENGTH];
    int totalCount = 0;

    // skip initial empty lines
    while (fgets(line, AOC24_DAY05_MAX_LINE_LENGTH, in) != NULL)
    {
        if (line[0] != '\n' && line[0] != '\r')
        {
            break;
        }
    }

    // parse ordering rules until blank line
    struct page_ordering_rule *page_ordering_rules = NULL;
    int rule_buffer_size = 0;
    if (parse_ordering_rules(in, &page_ordering_rules, &rule_buffer_size) != AOC_OK)
    {
        return AOC_ERR_PARSE;
    }

    int64_t page_ids[MAX_PAGE_IDS_PER_LINE];
    size_t page_ids_size;

    // indicator if we are between rules and page ID lines
    int in_rules = 0;
    // skip lines until blank line, then parse page ID lines until blank line or EOF
    while (fgets(line, AOC24_DAY05_MAX_LINE_LENGTH, in) != NULL)
    {
        if (line[0] == '\n' || line[0] == '\r')
        {
            if (!in_rules)
            {
                in_rules = 1;
                continue;
            }
            break;
        }
        page_ids_size = 0;

        if (aoc_util_parse_int_list_to_buf(line, ",", page_ids, MAX_PAGE_IDS_PER_LINE, &page_ids_size) != AOC_OK)
        {
            continue;
        }

        if (page_ids_size == 0)
        {
            continue;
        }

        int is_valid = 1;
        for (int i = 0; i < rule_buffer_size; ++i)
        {
            int before_idx = -1;
            int after_idx = -1;

            for (size_t j = 0; j < page_ids_size; ++j)
            {
                if (page_ids[j] == page_ordering_rules[i].before)
                {
                    before_idx = (int)j;
                }
                if (page_ids[j] == page_ordering_rules[i].after)
                {
                    after_idx = (int)j;
                }
            }

            if (before_idx != -1 && after_idx != -1)
            {
                if (before_idx > after_idx)
                {
                    is_valid = 0;
                    break;
                }
            }
        }

        if (is_valid)
        {
            totalCount += page_ids[page_ids_size / 2];
        }
    }
    free(page_ordering_rules);
    out_result->kind = AOC_RESULT_I64;
    out_result->value.i64 = totalCount;
    return AOC_OK;
}