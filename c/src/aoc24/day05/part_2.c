#include "aoc24/day05/day05.h"

#include "aoc_util/registry.h"
#include "aoc_util/parse_int_list.h"

aoc_error_t aoc24_day05_part2(FILE *in, aoc_result_t *out_result)
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

    // collect invalid rows in dynamic arrays so we can reorder them later
    int64_t **invalid_rows = NULL;
    size_t *invalid_lens = NULL;
    int invalid_count = 0;
    int invalid_capacity = 0;

    while (fgets(line, AOC24_DAY05_MAX_LINE_LENGTH, in) != NULL)
    {
        if (line[0] == '\n' || line[0] == '\r')
            break;

        page_ids_size = 0;
        if (aoc_util_parse_int_list_to_buf(line, ",", page_ids, MAX_PAGE_IDS_PER_LINE, &page_ids_size) != AOC_OK)
            continue;
        if (page_ids_size == 0)
            continue;

        // validate against ordering rules
        int is_valid = 1;
        for (int i = 0; i < rule_buffer_size && is_valid; ++i)
        {
            int before_idx = -1;
            int after_idx = -1;
            for (size_t j = 0; j < page_ids_size; ++j)
            {
                if (page_ids[j] == page_ordering_rules[i].before)
                    before_idx = (int)j;
                if (page_ids[j] == page_ordering_rules[i].after)
                    after_idx = (int)j;
            }
            if (before_idx != -1 && after_idx != -1 && before_idx > after_idx)
                is_valid = 0;
        }

        if (!is_valid)
        {
            if (invalid_count >= invalid_capacity)
            {
                invalid_capacity = invalid_capacity ? invalid_capacity * 2 : 16;
                invalid_rows = realloc(invalid_rows, sizeof(*invalid_rows) * invalid_capacity);
                invalid_lens = realloc(invalid_lens, sizeof(*invalid_lens) * invalid_capacity);
            }
            invalid_rows[invalid_count] = malloc(sizeof(int64_t) * page_ids_size);
            for (size_t k = 0; k < page_ids_size; ++k)
                invalid_rows[invalid_count][k] = page_ids[k];
            invalid_lens[invalid_count] = page_ids_size;
            invalid_count++;
        }
    }

    // helper to test whether rule a->b exists is done inline below

    for (int ir = 0; ir < invalid_count; ++ir)
    {
        int64_t *row = invalid_rows[ir];
        size_t len = invalid_lens[ir];
        int swapped;
        do
        {
            swapped = 0;
            for (size_t nidx = len - 1; nidx > 0; --nidx)
            {
                for (size_t i = 0; i < nidx; ++i)
                {
                    // check rule: row[nidx] before row[i]?
                    int can_before = 0;
                    for (int r = 0; r < rule_buffer_size; ++r)
                    {
                        if (page_ordering_rules[r].before == row[nidx] &&
                            page_ordering_rules[r].after == row[i])
                        {
                            can_before = 1;
                            break;
                        }
                    }
                    if (can_before)
                    {
                        int64_t tmp = row[nidx];
                        row[nidx] = row[i];
                        row[i] = tmp;
                        swapped = 1;
                    }
                }
            }
        } while (swapped);
        totalCount += row[len / 2];
        free(row);
    }
    free(invalid_rows);
    free(invalid_lens);
    free(page_ordering_rules);
    out_result->kind = AOC_RESULT_I64;
    out_result->value.i64 = totalCount;
    return AOC_OK;
}